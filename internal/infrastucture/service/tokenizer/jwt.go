package tokenizer

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ilfey/hikilist-go/internal/config/tokenizer"
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	repositoryInterface "github.com/ilfey/hikilist-go/internal/domain/repository/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"strconv"
	"time"
)

type JwtService struct {
	log   loggerInterface.Logger
	cfg   *tokenizer.Config
	token repositoryInterface.Token
}

func NewJwtService(
	log loggerInterface.Logger,
	cfg *tokenizer.Config,
	token repositoryInterface.Token,
) *JwtService {
	return &JwtService{
		log:   log,
		cfg:   cfg,
		token: token,
	}
}

func (s *JwtService) Generate(userId uint64) (*agg.TokenPair, error) {
	access, err := s.generateAccess(userId)
	if err != nil {
		return nil, err
	}

	refresh, err := s.generateRefresh(userId)
	if err != nil {
		return nil, err
	}

	return &agg.TokenPair{
		Access:  access,
		Refresh: refresh,
	}, nil
}

func (s *JwtService) generateAccess(userId uint64) (string, error) {
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(time.Duration(s.cfg.AccessLifeTime) * time.Hour)

	claims := jwt.MapClaims{
		"sub": strconv.FormatUint(userId, 10),
		"iss": s.cfg.Issuer,
		"iat": jwt.NewNumericDate(issuedAt),
		"exp": jwt.NewNumericDate(expiredAt),
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tkn.SignedString(s.cfg.Salt)
	if err != nil {
		return "", s.log.LogPropagate(err)
	}

	return token, nil
}

func (s *JwtService) generateRefresh(userId uint64) (string, error) {
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(time.Duration(s.cfg.RefreshLifeTime) * time.Hour)

	claims := jwt.MapClaims{
		"sub": strconv.FormatUint(userId, 10),
		"iss": s.cfg.Issuer,
		"iat": jwt.NewNumericDate(issuedAt),
		"exp": jwt.NewNumericDate(expiredAt),
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tkn.SignedString(s.cfg.Salt)
	if err != nil {
		return "", s.log.LogPropagate(err)
	}

	return token, nil
}

func (s *JwtService) Verify(ctx context.Context, token string) (uint64, error) {
	parsedToken, err := jwt.Parse(token, func(decodedToken *jwt.Token) (interface{}, error) {
		//if decodedToken.Header["alg"] != s.jwtTokenEncryptAlgo {
		//	// user must be banned here because the algo wasn't matched
		//	return nil, errors.New("invalid signing algorithm")
		//}

		// Cast to the configured givenToken signature type (stored in `s.jwtTokenEncryptAlgo`)
		_, success := decodedToken.Method.(*jwt.SigningMethodHMAC)
		if !success {
			return nil, errtype.NewTokenUnexpectedSigningMethodInternalError(token, decodedToken.Header["alg"])
		}

		// Salt is a string containing your secret, but you need pass the []byte
		return s.cfg.Salt, nil
	})
	if err != nil {
		// Parsing givenToken error occurred.
		s.log.Log(err)

		// Return a token invalid error.
		return 0, s.log.LogPropagate(errtype.NewAccessTokenIsInvalidError())
	}

	// Checking that token is not blocked
	found, err := s.token.Has(ctx, parsedToken.Raw)
	if err != nil {
		return 0, s.log.LogPropagate(err)
	}

	if found {
		return 0, s.log.LogPropagate(errtype.NewRefreshTokenWasBlockedError())
	}

	// Extract claims of the givenToken payload
	claims, success := parsedToken.Claims.(jwt.MapClaims)

	if success && parsedToken.Valid {
		err = s.isValidIssuer(token, claims)
		if err != nil {
			return 0, s.log.LogPropagate(errtype.NewAccessTokenIsInvalidError())
		}

		userID, err := s.getUserID(claims)
		if err != nil {
			return 0, s.log.LogPropagate(err)
		}

		return userID, nil
	}

	// Error occurred while extracting claims from givenToken or givenToken is not valid
	s.log.Log(errtype.NewTokenInvalidInternalError(token))

	return 0, s.log.LogPropagate(errtype.NewAccessTokenIsInvalidError())
}

func (s *JwtService) getUserID(claims jwt.Claims) (uint64, error) {
	// Extract subject (userId) from the claims
	stringId, err := claims.GetSubject()
	if err != nil {
		return 0, s.log.LogPropagate(err)
	}

	// Cast string to uint64
	userId, err := strconv.ParseUint(stringId, 10, 64)
	if err != nil {
		return 0, s.log.LogPropagate(err)
	}

	return userId, nil
}

func (s *JwtService) isValidIssuer(token string, claims jwt.Claims) error {
	// Extract the token issuer.
	iss, err := claims.GetIssuer()
	if err != nil {
		return s.log.LogPropagate(err)
	}

	// Check that token issuer is valid.
	if iss != s.cfg.Issuer {
		return s.log.LogPropagate(errtype.NewTokenIssuerWasNotMatchedInternalError(token))
	}

	return nil
}

func (s *JwtService) Block(ctx context.Context, tkn string) error {
	_, err := s.parseUserID(tkn)
	if err != nil {
		// block tkn anyway and log message,
		// because the tkn may be is invalid but must be blocked.
		// in this case a user will be undetermined

		s.log.Log(err)
	}

	createModel := &agg.TokenCreate{
		Token: tkn,
	}

	err = s.token.Create(ctx, createModel)
	if err != nil {
		return s.log.LogPropagate(err)
	}

	return nil
}

func (s *JwtService) parseUserID(token string) (uint64, error) {
	parsedToken, err := jwt.Parse(token, func(decodedToken *jwt.Token) (interface{}, error) {
		//if decodedToken.Header["alg"] != s.jwtTokenEncryptAlgo {
		//	// user must be banned here because the algo wasn't matched
		//	return nil, errors.New("invalid signing algorithm")
		//}

		// Cast to the configured givenToken signature type (stored in `s.jwtTokenEncryptAlgo`).
		_, success := decodedToken.Method.(*jwt.SigningMethodHMAC)
		if !success {
			return nil, errtype.NewTokenUnexpectedSigningMethodInternalError(token, decodedToken.Header["alg"])
		}

		// Salt is a string containing your secret, but you need pass the []byte.
		return s.cfg.Salt, nil
	})
	if err != nil {
		// Parsing givenToken error occurred.
		s.log.Log(err)

		// Return a token invalid error.
		return 0, s.log.LogPropagate(errtype.NewAccessTokenIsInvalidError())
	}

	claims, success := parsedToken.Claims.(jwt.MapClaims)
	if success {
		userID, err := s.getUserID(claims)
		if err != nil {
			return 0, s.log.LogPropagate(err)
		}

		return userID, nil
	}

	return 0, s.log.LogPropagate(errtype.NewAccessTokenIsInvalidError())
}
