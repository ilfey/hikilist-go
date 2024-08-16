package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/pkg/models/action"
	"github.com/ilfey/hikilist-go/pkg/repositories"
	"github.com/stretchr/testify/suite"
)

type ActionSuite struct {
	suite.Suite

	repo *repositories.ActionImpl
}

func (s *ActionSuite) SetupTest() {
	s.repo = &repositories.ActionImpl{}
}

func (s *ActionSuite) TestCreateSQL() {
	sql, args, err := s.repo.CreateSQL(&action.CreateModel{})
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"INSERT INTO user_actions (user_id,title,description,created_at) VALUES (?,?,?,?) RETURNING id",
		sql,
	)
}

func (s *ActionSuite) TestCountSQL() {
	sql, args, err := s.repo.CountSQL(nil)
	s.NoError(err)
	s.Nil(args)

	s.Equal(
		"SELECT COUNT(*) FROM user_actions",
		sql,
	)
}

func TestActionSuite(t *testing.T) {
	suite.Run(t, new(ActionSuite))
}
