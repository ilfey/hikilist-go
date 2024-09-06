package tests

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"testing"

	"github.com/ilfey/hikilist-go/internal/infrastucture/repositories"
	"github.com/stretchr/testify/suite"
)

type UserSuite struct {
	suite.Suite

	repo *repositories.User
}

func (s *UserSuite) SetupTest() {
	s.repo = &repositories.User{}
}

func (s *UserSuite) TestCreateSQL() {
	sql, args, err := s.repo.CreateSQL(&dto.UserCreateRequestDTO{})
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"INSERT INTO users (username,password) VALUES (?,?) RETURNING id",
		sql,
	)
}

func (s *UserSuite) TestGetSQL() {
	sql, args, err := s.repo.GetSQL(map[string]any{
		"id": 1,
	})
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"SELECT id, username, password, last_online, created_at FROM users WHERE id = ? LIMIT 1",
		sql,
	)
}

func (s *UserSuite) TestUpdatePasswordSQL() {
	sql, args, err := s.repo.UpdatePasswordSQL(1, "hash")
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"UPDATE users SET password = ? WHERE id = ?",
		sql,
	)
}

func (s *UserSuite) TestUpdateLastOnlineSQL() {
	sql, args, err := s.repo.UpdateLastOnlineSQL(1)
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"UPDATE users SET last_online = ? WHERE id = ?",
		sql,
	)
}

func (s *UserSuite) TestUpdateUsernameSQL() {
	sql, args, err := s.repo.UpdateUsernameSQL(1, "test")
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"UPDATE users SET username = ? WHERE id = ?",
		sql,
	)
}

func (s *UserSuite) TestDeleteSQL() {
	sql, args, err := s.repo.DeleteSQL(map[string]any{"id": 1})
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"DELETE FROM users WHERE id = ? RETURNING id",
		sql,
	)
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
