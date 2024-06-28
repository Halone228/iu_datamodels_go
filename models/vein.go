package iu_datamodels

import (
	"gorm.io/gorm"
)

type Vein struct {
	gorm.Model
	Name    string
	Url     string
	Slug    string
	Sources []*Source `gorm:"foreignKey:ID"`
}
