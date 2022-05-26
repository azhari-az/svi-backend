package service

import (
	"svi-backend/repository"
	"svi-backend/schema"
	"time"
)

type PostService interface {
	FindArticleById(Id int) (repository.Post, error)
	CreateArticle(postRequest schema.PostRequest) (repository.Post, error)
	UpdateArticle(postRequest schema.PostRequest, Id int) (repository.Post, error)
	DeleteArticle(Id int) (repository.Post, error)
	FindAllArticle(Page int, Limit int) ([]repository.Post, error)
}

type postService struct {
	repository repository.PostRepository
}

func PostNewService(repository repository.PostRepository) *postService {
	return &postService{repository}
}

func (s *postService) CreateArticle(postRequest schema.PostRequest) (repository.Post, error) {
	post := repository.Post{
		Title:       postRequest.Title,
		Content:     postRequest.Content,
		Category:    postRequest.Category,
		CreatedDate: time.Now(),
		Status:      postRequest.Status,
	}

	newPost, err := s.repository.CreateArticle(post)
	return newPost, err
}

func (s *postService) FindArticleById(ID int) (repository.Post, error) {
	return s.repository.FindArticleById(ID)
}

func (s *postService) UpdateArticle(postRequest schema.PostRequest, ID int) (repository.Post, error) {
	post := repository.Post{
		Title:       postRequest.Title,
		Content:     postRequest.Content,
		Category:    postRequest.Category,
		UpdatedDate: time.Now(),
		Status:      postRequest.Status,
	}
	newPost, err := s.repository.UpdateArticle(post, ID)
	return newPost, err
}

func (s *postService) DeleteArticle(ID int) (repository.Post, error) {
	return s.repository.DeleteArticle(ID)
}

func (s *postService) FindAllArticle(Page int, Limit int) ([]repository.Post, error) {
	return s.repository.FindAllArticle(Page, Limit)
}
