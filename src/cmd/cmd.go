package cmd

import (
	"net/http"
	"nft-marketplace-be/src/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func InitializeServer(cfg config.ENVConfig) {
	e := echo.New()

	v1 := e.Group("/v1")

	v1.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
	}))

	// Initialize Swagger
	docs.SwaggerInfo.Title = "NFT Marketplace API"
	docs.SwaggerInfo.Description = "This is NFT Marketplace API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
