package repository

import "time"

type PostRepository interface {
	FindArticleById(Id int) (Post, error)
	CreateArticle(post Post) (Post, error)
	DeleteArticle(Id int) (Post, error)
	UpdateArticle(post Post, Id int) (Post, error)
	FindAllArticle(Page int, Limit int) ([]Post, error)
}

type Post struct {
	Id          int32     `gorm:"column:Id; primaryKey; autoIncrement;"`
	Title       string    `gorm:"column:Title; type: varchar(200)"`
	Content     string    `gorm:"column:Content; type: text"`
	Category    string    `gorm:"column:Category; type: varchar(100)"`
	CreatedDate time.Time `gorm:"column:Created_date; type: timestamp; default:null"`
	UpdatedDate time.Time `gorm:"column:Updated_date; type: timestamp; default:null"`
	Status      string    `gorm:"column:Status; type: varchar(100)"`
}

func (r *DBRepository) CreateArticle(post Post) (Post, error) {
	err := r.db.Create(&post).Error

	return post, err
}

func (r *DBRepository) FindArticleById(ID int) (Post, error) {
	var post Post

	err := r.db.First(&post, ID).Error

	return post, err
}

func (r *DBRepository) UpdateArticle(p Post, ID int) (Post, error) {
	var post Post

	err := r.db.First(&post, ID).Updates(&p).Error

	return post, err
}

func (r *DBRepository) DeleteArticle(ID int) (Post, error) {
	var post Post

	err := r.db.First(&post, ID).Delete(ID).Error

	return post, err
}

func (r *DBRepository) FindAllArticle(Page int, Limit int) ([]Post, error) {
	var posts []Post

	err := r.db.Limit(Limit).Offset(Page - 1).Find(&posts).Error

	return posts, err
}
