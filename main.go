package main

import (
  "log"
  "time"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  "github.com/AnthonyLaiuppa/ezctf-api/db"
  Auth "github.com/AnthonyLaiuppa/ezctf-api/middleware"
  Controller "github.com/AnthonyLaiuppa/ezctf-api/controllers"
  jwt "gopkg.in/appleboy/gin-jwt.v2"
)

func main() {
  log.Println("Starting server..")

  db.Init();

  r := gin.Default()
  config := cors.DefaultConfig()
  config.AllowOrigins = []string{"http://localhost:3000"}
  r.Use(cors.New(config))

  jwtMiddleware := &jwt.GinJWTMiddleware{
    Realm:         "localhost",
    // store this somewhere, if your server restarts and you're
    // generating random passwords, any valid JWTs will be invalid
    Key:           []byte("NOTAREALSECRETTHISISJUSTAPLACEHOLDER"),
    Timeout:       time.Hour,
    MaxRefresh:    time.Hour * 24,
    Authenticator: Auth.Authenticate,
    // this method allows you to jump in and set user information
    // JWTs aren't encrypted, so don't store any sensitive info
    PayloadFunc:   Auth.Payload,
  }

  r.POST("/login", jwtMiddleware.LoginHandler)

  v1 := r.Group("/api/v1")
  {
    challenge := v1.Group("/challenge")
    {
      challenge.GET("/:id", Controller.GetChallenge)
      challenge.POST("/", Controller.CreateChallenge)
      challenge.PUT("/:id", Controller.UpdateChallenge)
      challenge.DELETE("/:id", Controller.DeleteChallenge)
    }
    challenges := v1.Group("/challenges")
    {
      challenges.GET("/", Controller.GetChallenges)
    }
    user := v1.Group("/user")
    {
      user.GET("/:id", Controller.GetUser)
    }
  }

  v1_auth := r.Group("/api/v1/secret/")
  v1_auth.Use(jwtMiddleware.MiddlewareFunc())
  {
    v1_auth.GET("/",Controller.AuthCheck)
  }


  r.Run()
}
