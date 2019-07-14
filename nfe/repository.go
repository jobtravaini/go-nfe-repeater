package nfe

import (
	"github.com/jinzhu/gorm"
)

type NfeRepository struct {
	Database *gorm.DB
}

func NewNfeRepository(db *gorm.DB) *NfeRepository {
	return &NfeRepository{Database: db}
}

func (repo NfeRepository) FindByAccessKey(accessKey string) (Nfe, error) {
	nfe := Nfe{}
	err := repo.Database.First(&nfe, Nfe{AccessKey: accessKey}).Error

	return nfe, err
}

func (repo NfeRepository) CreateOrUpdate(nfe Nfe) error {
	err := repo.Database.Where(Nfe{AccessKey: nfe.AccessKey}).Assign(Nfe{XmlValue: nfe.XmlValue}).FirstOrCreate(&nfe).Error

	return err
}
