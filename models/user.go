package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
	"strings"
)

// User "Object
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserName  string    `json:"user_name" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Score     int       `json:"score" binding:"required"`
	Solves    string    `json:"solves" binding:"required"`
	Elevated  bool      `json:"elevated" binding:"required" gorm:"default:'false'"`
}

//BeforeCreate ...
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	u := uuid.Must(uuid.NewV4())
	scope.SetColumn("ID", u.String())
	return nil
}

//BeforeUpdate ...
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

//HasntSolved ...
func (user *User) HasntSolved(id uuid.UUID) bool {

	//Check to avoid duplicate correct submission
	i := id.String()
	s := strings.Split(user.Solves, ",")
	return stringInSlice(i, s)

}

// stringInSlice ...
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return false
		}
	}
	return true
}