package controllers

import (
	"net/http"
	"nft-marketplace-be/src/config"
	"nft-marketplace-be/src/services"

	"github.com/labstack/echo/v4"
)

type NftController struct {
	envConfig  config.ENVConfig
	nftService services.NftServiceInterface
}

func NewNFTController(env config.ENVConfig, nftService services.NftServiceInterface) *NftController {
	return &NftController{
		envConfig:  env,
		nftService: nftService,
	}
}

func (o *NftController) BuildRoutes(e *echo.Group) {
	nfts := e.Group("/nfts")
	nfts.GET("/:address", o.GetNFTCollection)
}

// @Summary Get all projects
// @Description Get all projects
// @Tags NFT
// @Produce json
// @Router /v1/nfts/{address} [get]
// @Param address path string true "Contract Address"
func (o *NftController) GetNFTCollection(c echo.Context) error {
	address := c.Param("address")
	// Validate address

	collection, err := o.nftService.GetNFTCollection(address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, collection)
}
