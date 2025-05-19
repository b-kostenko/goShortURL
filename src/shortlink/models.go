package shortlink

import (
	"goShortURL/src/auth"
	"gorm.io/gorm"
)

type Shortlink struct {
	gorm.Model
	LongURL  string
	Slug     string `grom:"uniqueIndex"`
	ShortURL string
	UserID   uint
	User     auth.User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
