package services

import (
	"home/zellie/Code/guestbook-api/internal/interfaces"
	"home/zellie/Code/guestbook-api/internal/models"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type commentsService struct {
	log  zerolog.Logger
	repo interfaces.CommentsRepo
}

// creates an instance of CommentsService
func NewCommentsService(repo interfaces.CommentsRepo) (*commentsService, error) {
	return &commentsService{
		log: log.With().
			Str("component", "services.commentsService").
			Logger(),
		repo: repo,
	}, nil
}

func (cs commentsService) TestServiceFunc() error {
	cs.log.Trace().Msg("test comment services function was reached")
	return nil
}

func (cs commentsService) InsertComment(body *models.Comment) error {
	cs.log.Trace().Msg("attempting to post comment")

	if err := cs.validateBody(body); err != nil {
		cs.log.Error().
			Stack().
			Err(err).
			Msg("comment not posted")
		return err
	}

	cs.setCommentId(body)

	err := cs.repo.InsertComment(body)
	if err != nil {
		cs.log.Error().
			Stack().
			Err(err).
			Msg("unable to post comment")
		return err
	}

	cs.log.Debug().Msg("comment successfully posted")

	return nil
}

func (cs commentsService) GetAllComments() (*models.Comments, error) {
	cs.log.Trace().Msg("getting all comments")

	comments, err := cs.repo.GetAllComments()
	if err != nil {
		cs.log.Error().
			Stack().
			Err(err).
			Msg("unable to retrieve comments")

		return nil, err
	}

	cs.log.Debug().Msg("comments found")

	return comments, nil
}

func (cs commentsService) setCommentId(body *models.Comment) {
	body.CommentID = uuid.NewString()
}

func (cs commentsService) DeleteComment(cid string) error {
	cs.log.Trace().Msg("attempting to delete a comment")

	err := cs.repo.DeleteComment(cid)
	if err != nil {
		cs.log.Error().
			Stack().
			Err(err).
			Msg("unable to delete comment")
		return err
	}

	cs.log.Debug().Msg("comment successfully deleted")
	return nil
}

func (cs commentsService) validateBody(body *models.Comment) error {
	if body == nil || body.Description == "" {
		return models.BadRequestError("comment description required but was empty")
	}
	return nil
}
