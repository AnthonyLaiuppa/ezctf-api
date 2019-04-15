package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

// Challenge "Object
type Challenge struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"Name" binding:"required"`
	Category  string    `json:"Category" binding:"required"`
	Solves    int       `json:"Solves" binding:"required"`
	Points    int       `json:"Points" binding:"required"`
	Author    string    `json:"Author" binding:"required"`
	RawText   string    `json:"RawText" binding:"required"`
	Flag      string    `json:"Flag" binding:"required"`
}

//BeforeCreate ...
func (challenge *Challenge) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	u := uuid.Must(uuid.NewV4())
	scope.SetColumn("ID", u.String())
	return nil
}

//BeforeUpdate ...
func (challenge *Challenge) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
