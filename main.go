package main

import (
	Controller "github.com/AnthonyLaiuppa/ezctf-api/controllers"
	"github.com/AnthonyLaiuppa/ezctf-api/db"
	Auth "github.com/AnthonyLaiuppa/ezctf-api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	jwt "gopkg.in/appleboy/gin-jwt.v2"
	"log"
	"time"
)

func main() {
	log.Println("Starting server..")

	db.Init()

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	jwtMiddleware := &jwt.GinJWTMiddleware{
		Realm: "localhost",
		// store this somewhere, if your server restarts and you're
		// generating random passwords, any valid JWTs will be invalid
		Key:           []byte("NOTAREALSECRETTHISISJUSTAPLACEHOLDER"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour * 24,
		Authenticator: Auth.Authenticate,
		// this method allows you to jump in and set user information
		// JWTs aren't encrypted, so don't store any sensitive info
		PayloadFunc: Auth.Payload,
	}

	r.POST("/login", jwtMiddleware.LoginHandler)

	v1 := r.Group("/api/v1")
	{
		challenge := v1.Group("/challenge")
		{
			challenge.GET("/:id", Controller.GetChallenge)
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

	v1Auth := r.Group("/api/v1/secret/")
	v1Auth.Use(jwtMiddleware.MiddlewareFunc())
	{
		v1Auth.GET("/", Controller.AuthCheck)
	}

	modify := r.Group("/api/v1/modify/challenge")
	modify.Use(jwtMiddleware.MiddlewareFunc())
	{
		modify.POST("/", Controller.CreateChallenge)
		modify.PUT("/:id", Controller.UpdateChallenge)
		modify.DELETE("/:id", Controller.DeleteChallenge)
	}

	solve := r.Group("/api/v1/solve/challenge")
	solve.Use(jwtMiddleware.MiddlewareFunc())
	{
		solve.POST("/:id", Controller.SolveChallenge)
	}

	r.Run()
}
