package middleware

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/AnthonyLaiuppa/ezctf-api/db"
  "github.com/AnthonyLaiuppa/ezctf-api/models"
  jwt "gopkg.in/appleboy/gin-jwt.v2"
)


type Login struct {
	Email     string `form:"email" json:"Email" binding:"required"`
	Password string `form:"password" json:"Password"  binding:"required"`
}

func Payload(data interface{}) jwt.MapClaims{
	// in this method, you'd want to fetch some user info
	// based on their email address (which is provided once
	// they've successfully logged in).  the information
	// you set here will be available the lifetime of the
	// user's session
	if v, ok := data.(*models.User); ok {
		return jwt.MapClaims{
			"username": v.UserName,
			"elevated": v.Elevated,
		}
	}
	return jwt.MapClaims{}
}

func Authenticate(c *gin.Context) (interface{}, error) {
	var user models.User
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	db := db.GetDB()
 	if err := db.Where("Email = ?", json.Email).First(&user).Error; err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	c.ShouldBindJSON(&user)
	if json.Email != user.Email || json.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return "", jwt.ErrFailedAuthentication
	} else if json.Email == user.Email && json.Password == user.Password {
		return &user, nil
	}
	
	return "", jwt.ErrFailedAuthentication
}