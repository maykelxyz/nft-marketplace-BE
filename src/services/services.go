package services

import "nft-marketplace-be/src/repository"

type NftService struct {
	dbRepository repository.RepositoryInterface
}
type NftServiceInterface interface {
}

func NewNftService(repository repository.RepositoryInterface) NftServiceInterface {
	return &NftService{
		dbRepository: repository,
	}
}
