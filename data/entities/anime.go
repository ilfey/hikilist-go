package entities

import "gorm.io/gorm"

type Anime struct {
	gorm.Model

	Title            string
	Description      *string
	Poster           *string
	Episodes         *uint
	EpisodesReleased uint
	MalID            *uint
	ShikiID          *uint

	Related []*Anime `gorm:"many2many:animes_related;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Anime) TableName() string {
	return "animes"
}
