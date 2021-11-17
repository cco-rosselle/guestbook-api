package interfaces

import (
	"home/zellie/Code/guestbook-api/internal/models"
)

type CommentsRepo interface {
	InsertComment(c *models.Comment) error
	GetAllComments() (*models.Comments, error)
	DeleteComment(cid string) error
}
