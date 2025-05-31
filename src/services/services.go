package services

import "nft-marketplace-be/src/repository"

type NftService struct {
	dbRepository repository.RepositoryInterface
}
type NftServiceInterface interface {
	GetNFTCollection(address string) (bool, error)
}

func NewNftService(repository repository.RepositoryInterface) NftServiceInterface {
	return &NftService{
		dbRepository: repository,
	}
}

func (o NftService) GetNFTCollection(address string) (bool, error) {
	return true, nil
}
