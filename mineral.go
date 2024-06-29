package iu_datamodels_go

import (
	"gorm.io/gorm"
)

type attachment_id = string

type Attachment struct {
	ID attachment_id
}

type Mineral struct {
	gorm.Model
	HtmlText   string
	Source     *Source      `gorm:"foreignKey:ID"`
	Attachment []Attachment `gorm:"many2many:mineral_attachments;"`
	Tags       []*Tag       `gorm:"many2many:minerals_tags;"`
}
