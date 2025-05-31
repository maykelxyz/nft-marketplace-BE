package cmd

import (
	"net/http"
	"nft-marketplace-be/src/config"
	"nft-marketplace-be/src/controllers"
	"nft-marketplace-be/src/repository"
	"nft-marketplace-be/src/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitializeServer(cfg config.ENVConfig) {
	e := echo.New()

	v1 := e.Group("/v1")

	// Initialize Layers
	// Initialize Repository
	repository := repository.NewRepository(cfg)
	// Initialize Services
	nftService := services.NewNftService(repository)

	// Initialize Controllers
	nftController := controllers.NewNFTController(cfg, nftService)
	nftController.BuildRoutes(v1)
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

	// Routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
