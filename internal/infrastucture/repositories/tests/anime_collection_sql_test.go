package tests

import (
	"github.com/ilfey/hikilist-go/internal/domain/dto"
	"testing"

	"github.com/ilfey/hikilist-go/internal/infrastucture/repositories"
	"github.com/stretchr/testify/suite"
)

type AnimeCollectionSuite struct {
	suite.Suite

	repo *repositories.AnimeCollection
}

func (s *AnimeCollectionSuite) SetupTest() {
	s.repo = &repositories.AnimeCollection{}
}

func (s *AnimeCollectionSuite) TestGetCollectionIdSQL() {
	sql, args, err := s.repo.GetCollectionIdSQL(1, 1)
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"SELECT id FROM collections WHERE id = ? AND user_id = ?",
		sql,
	)
}

//goland:noinspection ALL
func (s *AnimeCollectionSuite) TestAddAnimesSQL() {
	// Test one anime.
	aam := &dto.CollectionAddAnimeRequestDTO{
		UserID:       1,
		CollectionID: 1,

		Animes: []uint64{1},
	}

	sql, args, err := s.repo.AddAnimesSQL(aam)
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"INSERT INTO animes_collections (collection_id,anime_id) VALUES (?,?)",
		sql,
	)

	// Test two animes.
	aam = &dto.CollectionAddAnimeRequestDTO{
		UserID:       1,
		CollectionID: 1,

		Animes: []uint64{1, 2},
	}

	sql, args, err = s.repo.AddAnimesSQL(aam)
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"INSERT INTO animes_collections (collection_id,anime_id) VALUES (?,?),(?,?)",
		sql,
	)

	// Test panic when animes is empty.
	s.Panics(func() {
		aam := &dto.CollectionAddAnimeRequestDTO{
			UserID:       1,
			CollectionID: 1,

			Animes: []uint64{},
		}

		s.repo.AddAnimesSQL(aam)
	})
}

//goland:noinspection ALL
func (s *AnimeCollectionSuite) TestRemoveAnimesSQL() {
	// Test one anime.
	ram := &dto.CollectionRemoveAnimeRequestDTO{
		UserID:       1,
		CollectionID: 1,

		Animes: []uint64{1},
	}

	sql, args, err := s.repo.RemoveAnimesSQL(ram)
	s.NoError(err)
	s.Nil(args)

	s.Equal(
		"DELETE FROM animes_collections WHERE collection_id = 1 AND anime_id IN (1)",
		sql,
	)

	// Test two animes.
	ram = &dto.CollectionRemoveAnimeRequestDTO{
		UserID:       1,
		CollectionID: 1,

		Animes: []uint64{1, 2},
	}

	sql, args, err = s.repo.RemoveAnimesSQL(ram)
	s.NoError(err)
	s.Nil(args)

	s.Equal(
		"DELETE FROM animes_collections WHERE collection_id = 1 AND anime_id IN (1,2)",
		sql,
	)

	// Test panic when animes is empty.
	s.Panics(func() {
		ram := &dto.CollectionRemoveAnimeRequestDTO{
			UserID:       1,
			CollectionID: 1,

			Animes: []uint64{},
		}

		s.repo.RemoveAnimesSQL(ram)
	})
}

func TestAnimeCollectionSuite(t *testing.T) {
	suite.Run(t, new(AnimeCollectionSuite))
}
