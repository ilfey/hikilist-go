package entities

import "gorm.io/gorm"

type Anime struct {
	gorm.Model

	// Название аниме
	Title string
	
	// Описание
	Description *string

	// Постер
	Poster *string

	// Количество эпизодов
	Episodes *uint

	// Количество вышедших эпизодов
	EpisodesReleased uint

	// ID в MAL
	//
	// Если `nil` - то аниме не имеет MAL ID.
	// Значение должно быть уникальным, для возможности поиска по этому полю.
	MalID *uint `gorm:"unique"`

	// ID в shikimori
	//
	// Если `nil` - то аниме не имеет shikimori ID.
	// Значение должно быть уникальным, для возможности поиска по этому полю.
	ShikiID *uint `gorm:"unique"`

	// Связанные аниме
	Related []*Anime `gorm:"many2many:animes_related;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Anime) TableName() string {
	return "animes"
}
