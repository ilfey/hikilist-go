package tests

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"testing"

	"github.com/ilfey/hikilist-go/internal/infrastucture/repositories"
	"github.com/stretchr/testify/suite"
)

type AnimeSuite struct {
	suite.Suite

	repo *repositories.Anime
}

func (s *AnimeSuite) SetupTest() {
	s.repo = &repositories.Anime{}
}

func (s *AnimeSuite) TestCreateSQL() {
	sql, args, err := s.repo.CreateSQL(&dto.AnimeCreateRequestDTO{})
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"INSERT INTO animes (title,description,poster,episodes,episodes_released,mal_id,shiki_id) VALUES (?,?,?,?,?,?,?) RETURNING id",
		sql,
	)
}

func (s *AnimeSuite) TestGetSQL() {
	sql, args, err := s.repo.GetSQL(map[string]any{
		"id": 1,
	})
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"SELECT id, title, description, poster, episodes, episodes_released, mal_id, shiki_id, created_at, updated_at FROM animes WHERE id = ? LIMIT 1",
		sql,
	)
}

//
//func (s *AnimeSuite) TestFindWithPaginatorSQL() {
//	sql, args, err := s.repo.FindWithPaginatorSQL(builder.NewAnimePaginator(map[string][]string{}), nil)
//	s.NoError(err)
//	s.Nil(args)
//
//	s.Equal(
//		"SELECT id, title, poster, episodes, episodes_released "+
//			"FROM animes ORDER BY id DESC LIMIT 10 OFFSET 0",
//		sql,
//	)
//}

//func (s *AnimeSuite) TestFindFromCollectionWithPaginatorSQL() {
//	sql, args, err := s.repo.FindFromCollectionWithPaginatorSQL(builder.NewAnimePaginator(map[string][]string{}), 1, 1)
//	s.NoError(err)
//	s.NotNil(args)
//
//	s.Equal(
//		"SELECT id, title, poster, episodes, episodes_released FROM animes_collections "+
//			"JOIN animes ON animes.id = animes_collections.anime_id "+
//			"WHERE collection_id = (SELECT id FROM collections WHERE id = ? AND (is_public = TRUE OR user_id = ?)) "+
//			"ORDER BY id DESC LIMIT 10 OFFSET 0",
//		sql,
//	)
//}

func (s *AnimeSuite) TestCountSQL() {
	sql, args, err := s.repo.CountSQL(nil)
	s.NoError(err)
	s.Nil(args)

	s.Equal(
		"SELECT COUNT(*) FROM animes",
		sql,
	)
}

func (s *AnimeSuite) TestCountInCollectionSQL() {
	req := &dto.AnimeListFromCollectionRequestDTO{
		CollectionID: 1,
		UserID:       1,
	}

	sql, args, err := s.repo.CountInCollectionSQL(req)
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"SELECT COUNT(*) FROM animes_collections "+
			"JOIN collections ON collections.id = animes_collections.collection_id "+
			"WHERE collection_id = ? AND (is_public = TRUE OR user_id = ?)",
		sql,
	)
}

func TestAnimeSuite(t *testing.T) {
	suite.Run(t, new(AnimeSuite))
}
