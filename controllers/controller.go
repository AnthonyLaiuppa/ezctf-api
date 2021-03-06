package controller

import (
	"github.com/AnthonyLaiuppa/ezctf-api/db"
	"github.com/AnthonyLaiuppa/ezctf-api/models"
	"github.com/gin-gonic/gin"
	jwt "gopkg.in/appleboy/gin-jwt.v2"
	"net/http"
	"strings"
)

//GetChallenge ...
func GetChallenge(c *gin.Context) {

	id := c.Param("id")
	var challenge models.Challenge
	db := db.GetDB()
	if err := db.Where("id = ?", id).First(&challenge).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(200, challenge)
}

//GetChallenges ...
func GetChallenges(c *gin.Context) {

	var challenges []models.Challenge
	db := db.GetDB()
	db.Limit(25).Find(&challenges)
	c.JSON(200, challenges)
}

//CreateChallenge ...
func CreateChallenge(c *gin.Context) {
	var challenge models.Challenge
	var db = db.GetDB()

	if err := c.BindJSON(&challenge); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Create(&challenge)
	c.JSON(http.StatusOK, &challenge)
}

//UpdateChallenge ...
func UpdateChallenge(c *gin.Context) {
	id := c.Param("id")
	var challenge models.Challenge

	db := db.GetDB()
	if err := db.Where("id = ?", id).First(&challenge).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&challenge)
	db.Save(&challenge)
	c.JSON(http.StatusOK, &challenge)
}

//DeleteChallenge ...
func DeleteChallenge(c *gin.Context) {
	id := c.Param("id")
	var challenge models.Challenge
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&challenge).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&challenge)
}

//GetUser ...
func GetUser(c *gin.Context) {

	id := c.Param("id")
	var user models.User
	db := db.GetDB()
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(200, user)
}

//AuthCheck ...
func AuthCheck(c *gin.Context) {
	// the JWT middleware provides a useful method to extract
	// custom claims, it's basically the reverse of what's being
	// done in the payload function below
	claims := jwt.ExtractClaims(c)

	// for this example, we'll just dump out our custom claims
	// but in reality you could create your own middleware
	// handler to intercept this information and provide an
	// additional level of role-based security
	c.String(http.StatusOK, "id: %s\nrole: %s", claims["username"], claims["elevated"])
}

//Solve ...
type Solve struct {
	Flag string `form:"flag" json:"Flag" binding:"required"`
}

//SolveChallenge ...
func SolveChallenge(c *gin.Context) {

	//First check that JSON received binds
	var s Solve
	if err := c.BindJSON(&s); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//Next grab information on challenge and user
	claims := jwt.ExtractClaims(c)
	id := c.Param("id")
	un := claims["username"]
	var user models.User
	var challenge models.Challenge
	db := db.GetDB()

	if err := db.Where("user_name = ?", un).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := db.Where("id = ?", id).First(&challenge).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}


	uhs := user.HasntSolved(challenge.ID)
	//Make sure flags match and user hasnt already solved
	if s.Flag == challenge.Flag && uhs == true {

		user.Score = user.Score + challenge.Points
		user.Solves = strings.Join([]string{user.Solves, challenge.ID.String()}, ",") 
		challenge.Solves++

		db.Save(&challenge)
		db.Save(&user)
		c.JSON(http.StatusOK, &challenge)
		return
	}else if uhs == false {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.AbortWithStatus(http.StatusNotAcceptable )
	return
}


