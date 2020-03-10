package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Content 		string
	Title   		string
	Excerpt			string
	Categorys		[]*Category `gorm:"many2many:article_categorys;"`
	UserArticles	UserArticles 	`gorm:"foreignkey:ArticleId"`
}