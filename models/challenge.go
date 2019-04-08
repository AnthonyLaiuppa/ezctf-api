package models

import (
  "time"
  "github.com/satori/go.uuid"
  "github.com/jinzhu/gorm"
)

// Challenge "Object
type Challenge struct {
  ID  uuid.UUID   `json:"id"`
  CreatedAt time.Time  `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  Name string  `json:"Name" binding:"required"`
  Category string  `json:"Category" binding:"required"`
  Solves int  `json:"Solves" binding:"required"`
  Points int  `json:"Points" binding:"required"`
  Author string  `json:"Author" binding:"required"`
  RawText string  `json:"RawText" binding:"required"`
}

func (challenge *Challenge) BeforeCreate(scope *gorm.Scope) error {
  scope.SetColumn("CreatedAt", time.Now())
  u := uuid.Must(uuid.NewV4())
  scope.SetColumn("ID", u.String())
  return nil
}

func (challenge *Challenge) BeforeUpdate(scope *gorm.Scope) error {
  scope.SetColumn("UpdatedAt", time.Now())
  return nil
}
