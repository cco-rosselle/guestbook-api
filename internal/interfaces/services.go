package interfaces

import (
	"home/zellie/Code/guestbook-api/internal/models"
)

type CommentsService interface {
	TestServiceFunc() error
	PostComment(*models.Comment) error
	GetAllComments() (*models.Comments, error)
}