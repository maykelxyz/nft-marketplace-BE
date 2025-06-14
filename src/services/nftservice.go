package services

import "nft-marketplace-be/src/repository"

type NftService struct {
	dbRepository repository.RepositoryInterface
}
type NftServiceInterface interface {
	ListFromAddress(address string) (bool, error)
	Count(address string) (int, error)
}

func NewNftService(repository repository.RepositoryInterface) NftServiceInterface {
	return &NftService{
		dbRepository: repository,
	}
}

func (o NftService) ListFromAddress(address string) (bool, error) {
	return true, nil
}

func (o NftService) Count(address string) (int, error) {
	return 0, nil
}
