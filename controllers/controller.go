package controller

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/AnthonyLaiuppa/ezctf-api/models"
  "github.com/AnthonyLaiuppa/ezctf-api/db"
  jwt "gopkg.in/appleboy/gin-jwt.v2"
)

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

func GetChallenges(c *gin.Context) {

  var challenges []models.Challenge
  db := db.GetDB()
  db.Limit(25).Find(&challenges)
  c.JSON(200, challenges)
}

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