package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	"github.com/rafaeljesus/kyp-auth/db"
	"github.com/rafaeljesus/kyp-auth/handlers"
	"github.com/rafaeljesus/kyp-auth/models"
	"log"
	"os"
)

func main() {
	db.Connect()
	db.Repo.AutoMigrate(&models.User{})

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

	v1 := e.Group("/v1")
	v1.GET("/healthz", handlers.HealthzIndex)
	v1.POST("/users", handlers.UsersCreate)
	v1.POST("/token", handlers.TokenCreate)

	log.Print("Starting Kyp Auth Service...")

	e.Run(fasthttp.New(":" + os.Getenv("KYP_AUTH_PORT")))
}
