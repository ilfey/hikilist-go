package tests

import (
	"testing"

	animecollection "github.com/ilfey/hikilist-go/pkg/models/anime_collection"
	"github.com/ilfey/hikilist-go/pkg/repositories"
	"github.com/stretchr/testify/suite"
)

type AnimeCollectionSuite struct {
	suite.Suite

	repo *repositories.AnimeCollectionImpl
}

func (s *AnimeCollectionSuite) SetupTest() {
	s.repo = &repositories.AnimeCollectionImpl{}
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

func (s *AnimeCollectionSuite) TestAddAnimesSQL() {
	// Test one anime.
	aam := &animecollection.AddAnimesModel{
		UserID:       1,
		CollectionID: 1,

		Animes: []uint{1},
	}

	sql, args, err := s.repo.AddAnimesSQL(aam)
	s.NoError(err)
	s.NotNil(args)

	s.Equal(
		"INSERT INTO animes_collections (collection_id,anime_id) VALUES (?,?)",
		sql,
	)

	// Test two animes.
	aam = &animecollection.AddAnimesModel{
		UserID:       1,
		CollectionID: 1,

		Animes: []uint{1, 2},
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
		aam := &animecollection.AddAnimesModel{
			UserID:       1,
			CollectionID: 1,

			Animes: []uint{},
		}

		s.repo.AddAnimesSQL(aam)
	})
}

func (s *AnimeCollectionSuite) TestRemoveAnimesSQL() {
	// Test one anime.
	ram := &animecollection.RemoveAnimesModel{
		UserID:       1,
		CollectionID: 1,

		Animes: []uint{1},
	}

	sql, args, err := s.repo.RemoveAnimesSQL(ram)
	s.NoError(err)
	s.Nil(args)

	s.Equal(
		"DELETE FROM animes_collections WHERE collection_id = 1 AND anime_id IN (1)",
		sql,
	)

	// Test two animes.
	ram = &animecollection.RemoveAnimesModel{
		UserID:       1,
		CollectionID: 1,

		Animes: []uint{1, 2},
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
		ram := &animecollection.RemoveAnimesModel{
			UserID:       1,
			CollectionID: 1,

			Animes: []uint{},
		}

		s.repo.RemoveAnimesSQL(ram)
	})
}

func TestAnimeCollectionSuite(t *testing.T) {
	suite.Run(t, new(AnimeCollectionSuite))
}
