package model

import (
	"time"
)

type Pokemon struct {
	ID         uint `gorm:"primaryKey"`
	No         uint
	Name       string
	Species    string
	H          uint
	A          uint
	B          uint
	C          uint
	D          uint
	S          uint
	CreatedOn  time.Time `gorm:"autoCreateTime"`
	ModifiedOn time.Time `gorm:"autoUpdateTime:true"`
}
