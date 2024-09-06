package tests

import (
	"github.com/ilfey/hikilist-go/internal/domain/agg"
	"testing"

	"github.com/ilfey/hikilist-go/internal/infrastucture/repositories"
	"github.com/stretchr/testify/suite"
)

type TokenSuite struct {
	suite.Suite

	repo *repositories.Token
}

func (s *TokenSuite) SetupTest() {
	s.repo = &repositories.Token{}
}

func (s *TokenSuite) TestCreateSQL() {
	sql, args, err := s.repo.CreateSQL(&agg.TokenCreate{})
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"INSERT INTO tokens (token) VALUES (?) RETURNING id",
		sql,
	)
}

func (s *TokenSuite) TestGetSQL() {
	sql, args, err := s.repo.GetSQL(map[string]any{
		"id": 1,
	})
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"SELECT id, token, created_at FROM tokens WHERE id = ? LIMIT 1",
		sql,
	)
}

func (s *TokenSuite) TestDeleteSQL() {
	sql, args, err := s.repo.DeleteSQL(map[string]any{"id": 1})
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"DELETE FROM tokens WHERE id = ? RETURNING id",
		sql,
	)
}

func TestTokenSuite(t *testing.T) {
	suite.Run(t, new(TokenSuite))
}
