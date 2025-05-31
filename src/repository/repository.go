package repository

import (
	"nft-marketplace-be/src/config"
)

type Repository struct {
	EnvConfig config.ENVConfig
}

type RepositoryInterface interface{}

func NewRepository(cfg config.ENVConfig) RepositoryInterface {
	return &Repository{
		EnvConfig: cfg,
	}
}
