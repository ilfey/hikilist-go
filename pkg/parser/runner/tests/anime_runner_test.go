package tests

import (
	"context"
	"testing"

	"github.com/ilfey/hikilist-go/pkg/parser/mocks/extractor"
	"github.com/ilfey/hikilist-go/pkg/parser/models"
	"github.com/ilfey/hikilist-go/pkg/parser/runner"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAnimeRunnerTick(t *testing.T) {
	type args struct {
		pages   []*models.AnimeListModel
		details []*models.AnimeDetailModel
	}

	testCases := []struct {
		desc    string
		wantErr bool
		args    args
	}{
		{
			desc:    "Tick with empty page",
			wantErr: true,
			args: args{
				pages: []*models.AnimeListModel{},
			},
		},
		{
			desc:    "Tick with one page",
			wantErr: true,
			args: args{
				pages: []*models.AnimeListModel{
					{IDs: []uint{1}},
				},
				details: []*models.AnimeDetailModel{
					{
						ID: 1,
					},
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			animeExtractor := extractor.NewAnimeExtractor(t)

			animeExtractor.
				On("FetchList", mock.AnythingOfType("uint")).
				Times(len(tC.args.pages)+1). // +1 for empty page.
				Return(func(page uint) *models.AnimeListModel {
					// Return empty page if out of range.
					if page >= uint(len(tC.args.pages)) {
						return &models.AnimeListModel{}
					}

					// Return page.
					return tC.args.pages[page]
				}, nil)

			if len(tC.args.details) != 0 {
				animeExtractor.
					On("FetchById", mock.AnythingOfType("uint")).
					Times(len(tC.args.details)).
					Return(func(id uint) *models.AnimeDetailModel {
						for _, d := range tC.args.details {
							if d.ID == id {
								return d
							}
						}

						return nil
					}, nil)

				animeExtractor.
					On("Resolve", mock.Anything, mock.Anything).
					Times(len(tC.args.details)).
					Return(nil)
			}

			stats := runner.NewStatistic()

			animeRunner := runner.NewAnimeRunner(
				logrus.New(),
				animeExtractor,
				stats,
			)

			for i := 0; i < len(tC.args.pages)+1; i++ {

				err := animeRunner.Tick(context.Background())
				// If page is empty
				if i == len(tC.args.pages) {
					if !tC.wantErr {
						assert.NoError(t, err)
					} else {
						assert.Error(t, err)
					}
				} else {
					assert.NoError(t, err)
				}

				stats.AddPages(1)
			}
		})
	}
}