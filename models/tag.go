package iu_datamodels

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Descriptor string
	Minerals   []*Mineral `gorm:"many2many:mineral_tags;"`
}
