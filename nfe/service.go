package nfe

import (
	"log"
)

type NfeService struct {
	Repository INfeRepository
}

func NewNfeService(repository INfeRepository) *NfeService {
	return &NfeService{Repository:repository}
}

func (service NfeService) GetNfe(accessKey string) (Nfe, error) {
	nfe, err := service.Repository.FindByAccessKey(accessKey)

	if err != nil {
		log.Printf("Error while retrieving nfe access key %s: %s\n", accessKey, err)
	}

	return nfe, err
}
