package main

import (
	"context"

	"github.com/go-kit/kit/log"
)

type Author struct {
	Name        string `json:"name,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Birth       string `json:"birth,omitempty"`
	Nationality string `json:"Nationality,omitempty"`
}
type authorservice struct {
	logger log.Logger
}

// Service describes the Author service.
type AuthorService interface {
	CreateAuthor(ctx context.Context, author Author) (string, error)
	GetAuthorById(ctx context.Context, id string) (interface{}, error)
	UpdateAuthor(ctx context.Context, author Author) (string, error)
	DeleteAuthor(ctx context.Context, id string) (string, error)
}

var authors = []Author{
	Author{
		Name:        "John Doe",
		Gender:      "N/D",
		Birth:       "11-02-1991",
		Nationality: "USA"},
}

func findAuthor(x string) int {
	for i, author := range authors {
		if x == author.Name {
			return i
		}
	}
	return -1
}

func NewService(logger log.Logger) AuthorService {
	return &authorservice{
		logger: logger,
	}
}

func (s authorservice) CreateAuthor(ctx context.Context, author Author) (string, error) {
	var msg = "success"
	authors = append(authors, author)
	return msg, nil
}

func (s authorservice) GetAuthorById(ctx context.Context, id string) (interface{}, error) {
	var err error
	var author interface{}
	var empty interface{}
	i := findAuthor(id)
	if i == -1 {
		return empty, err
	}
	author = authors[i]
	return author, nil
}
func (s authorservice) DeleteAuthor(ctx context.Context, id string) (string, error) {
	var err error
	msg := ""
	i := findAuthor(id)
	if i == -1 {
		return "", err
	}
	copy(authors[i:], authors[i+1:])
	authors[len(authors)-1] = Author{}
	authors = authors[:len(authors)-1]
	return msg, nil
}
func (s authorservice) UpdateAuthor(ctx context.Context, Author Author) (string, error) {
	var empty = ""
	var err error
	var msg = "success"
	i := findAuthor(Author.Name)
	if i == -1 {
		return empty, err
	}
	authors[i] = Author
	return msg, nil
}
