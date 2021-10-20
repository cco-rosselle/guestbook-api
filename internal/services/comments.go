package services

import (
	"home/zellie/Code/guestbook-api/internal/models"
	"home/zellie/Code/guestbook-api/internal/repos"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	
)

type CommentsService struct {
	log zerolog.Logger
	repo *repos.CommentsRepo
}

// creates an instance of CommentsService 
func NewCommentsService(repo *repos.CommentsRepo) *CommentsService {
	return &CommentsService {
		log: log.With().
			Str("component", "services.commentsService").
			Logger(),
		repo: repo,
	}
}


func (cs CommentsService) TestServiceFunc() error {
	cs.log.Trace().Msg("test comment services function was reached")
	return nil
}

func (cs CommentsService) PostComment(body *models.Comment) error {
	cs.log.Trace().Msg("attempting to post comment")

	// validate if there's a description

	cs.setCommentId(body)

	if err := cs.repo.InsertComment(body); err != nil {
		cs.log.Error().
			Stack().
			Err(err).
			Msg("unable to post comment")
		return err
	}

	cs.log.Debug().Msg("comment successfully posted")

	return nil
}

func (cs CommentsService) GetAllComments() (*models.Comments, error) {
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

func (cs CommentsService) setCommentId(body *models.Comment) {
	body.CommentID = uuid.NewString()
}

