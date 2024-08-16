package tests

import (
	"testing"

	"github.com/ilfey/hikilist-go/pkg/models/collection"
	"github.com/ilfey/hikilist-go/pkg/repositories"
	"github.com/stretchr/testify/suite"
)

type CollectionSuite struct {
	suite.Suite

	repo *repositories.CollectionImpl
}

func (s *CollectionSuite) SetupTest() {
	s.repo = &repositories.CollectionImpl{}
}

func (s *CollectionSuite) TestCreateSQL() {
	sql, args, err := s.repo.CreateSQL(&collection.CreateModel{})
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"INSERT INTO collections (title,user_id,description,is_public) VALUES (?,?,?,?) RETURNING id",
		sql,
	)
}

func (s *CollectionSuite) TestGetSQL() {
	sql, args, err := s.repo.GetSQL(map[string]any{"id": 1})
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"SELECT id, title, user_id, description, is_public, created_at, updated_at FROM collections WHERE id = ? LIMIT 1",
		sql,
	)
}

func (s *CollectionSuite) TestFindSQL() {
	sql, args, err := s.repo.FindSQL(collection.NewPaginatorFromQuery(map[string][]string{}), nil)
	s.NoError(err)
	s.Nil(args)

	s.Equal(
		"SELECT id, user_id, title, description FROM collections ORDER BY id DESC LIMIT 10 OFFSET 0",
		sql,
	)
}

func (s *CollectionSuite) TestCountSQL() {
	sql, args, err := s.repo.CountSQL(nil)
	s.NoError(err)
	s.Nil(args)

	s.Equal(
		"SELECT COUNT(*) FROM collections",
		sql,
	)
}

func TestCollectionSuite(t *testing.T) {
	suite.Run(t, new(CollectionSuite))
}
