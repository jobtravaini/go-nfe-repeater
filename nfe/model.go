package nfe

import(
	"github.com/jinzhu/gorm"
)

type Nfe struct {
	gorm.Model
	AccessKey string `gorm:"type:varchar(255);unique_index"`
	XmlValue string `gorm:"type:clob"`
}