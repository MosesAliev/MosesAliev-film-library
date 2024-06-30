package models

import (
	"time"

	"gorm.io/gorm"
)

type Actor struct {
	gorm.Model
	Name string    `json:"name"`
	Sex  bool      `json:"sex"`
	Born time.Time `json:"born"`
}
