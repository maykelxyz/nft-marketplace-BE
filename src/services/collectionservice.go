package services

import (
	"nft-marketplace-be/src/config"
	"nft-marketplace-be/src/models"
)

type CollectionServiceInterface interface {
	Count(address string) (int, error)
	ListFromAddress(address string) ([]models.Nft, error)
}

type CollectionService struct {
	envConfig config.ENVConfig
}

func NewCollectionService(envConfig config.ENVConfig) CollectionServiceInterface {
	return &CollectionService{
		envConfig: envConfig,
	}
}

func (s *CollectionService) Count(address string) (int, error) {
	count, err := s.envConfig.DB.Model(&models.Nft{}).Where("collection_address = ?", address).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *CollectionService) ListFromAddress(address string) ([]models.Nft, error) {
	var nfts []models.Nft
	err := s.envConfig.DB.Model(&models.Nft{}).Where("collection_address = ?", address).Find(&nfts).Error
	if err != nil {
		return nil, err
	}
	return nfts, nil
}
