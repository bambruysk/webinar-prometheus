package models

import (
	"blogpost/internal/models"
)

type CreateBlogPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type BlogPosGetResponse struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID string `json:"author_id"`
}

func ConvertToBlogPost(post *CreateBlogPost) *models.BlogPost {
	return &models.BlogPost{
		Title:    post.Title,
		Content:  post.Content,
		AuthorID: post.Author,
	}
}

func ConvertToGetBlogPost(post *models.BlogPost) *BlogPosGetResponse {
	return &BlogPosGetResponse{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		AuthorID: post.AuthorID,
	}
}
