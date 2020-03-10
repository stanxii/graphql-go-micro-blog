package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string
	Articles []*Article `gorm:"many2many:article_categorys;"`
}