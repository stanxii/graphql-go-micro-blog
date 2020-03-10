package model

//type User struct {
//	ID        uint `gorm:"primary_key"`
//	Articles []*Article `gorm:"many2many:user_articles;"`
//}

type UserArticles struct {
	ArticleId	uint `gorm:"primary_key"`
	UserId		uint `gorm:"primary_key"`
}