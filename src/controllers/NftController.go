package controllers

import (
	"nft-marketplace-be/src/config"
	"nft-marketplace-be/src/services"

	"github.com/labstack/echo/v4"
)

type NftController struct {
	ENVConfig  config.ENVConfig
	NftService services.NftServiceInterface
}

func NewNFTController(env config.ENVConfig, nftService services.NftServiceInterface) *NftController {
	return &NftController{
		ENVConfig:  env,
		NftService: nftService,
	}
}

func (o *NftController) BuildRoutes(e *echo.Group) {}
