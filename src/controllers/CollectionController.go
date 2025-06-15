package controllers

import (
	"nft-marketplace-be/src/config"
	"nft-marketplace-be/src/services"
)

type CollectionController struct {
	envConfig         config.ENVConfig
	collectionService services.CollectionServiceInterface
}

func NewCollectionController(envConfig config.ENVConfig, collectionService services.CollectionServiceInterface) *CollectionController {
	return &CollectionController{
		envConfig: envConfig,
	}
}
