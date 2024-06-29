package iu_datamodels_go

import (
	"gorm.io/gorm"
)

type Source struct {
	gorm.Model
	Slug     string     `gorm:"unique"`
	Vein     *Vein      `gorm:"foreignKey:ID"`
	Minerals []*Mineral `gorm:"foreignKey:ID"`
}
