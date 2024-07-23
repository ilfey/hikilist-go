package entities

import "gorm.io/gorm"

type Collection struct {
	gorm.Model

	UserID uint
	User   *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:SET NULL"`

	Name string `gorm:"not null"`

	Description *string

	IsPublic bool `gorm:"not null"`

	Animes []*Anime `gorm:"many2many:animes_collections"`
}

func (Collection) TableName() string {
	return "collections"
}
