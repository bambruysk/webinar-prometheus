package models

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrorPostNotFound  = errors.New("post not found")
	ErrorAlreadyExists = errors.New("post already exists")
)

type BlogPost struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID string `json:"author_id"`
}

func NewID() string {
	return uuid.Must(uuid.NewRandom()).String()
}
