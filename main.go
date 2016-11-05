package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	"github.com/rafaeljesus/kyp-auth/handlers"
	"log"
	"os"
)

var KYP_AUTH_PORT = os.Getenv("KYP_AUTH_PORT")

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.Gzip())

	v1 := e.Group("/v1")
	v1.GET("/healthz", handlers.HealthzIndex)
	v1.POST("/token", handlers.TokenCreate)

	log.Print("Starting Kyp Auth Service at port " + KYP_AUTH_PORT)

	e.Run(fasthttp.New(":" + KYP_AUTH_PORT))
}
