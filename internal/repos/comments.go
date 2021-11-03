package repos

import (
	"home/zellie/Code/guestbook-api/internal/interfaces"
	"home/zellie/Code/guestbook-api/internal/models"

	"github.com/rs/zerolog"
	_ "github.com/lib/pq"
)


type commentsRepo struct {
	log zerolog.Logger
	pc *PostgresRepo
}

func NewCommentsRepo() (interfaces.CommentsRepo, error) {
	repo, err := newPostgresRepo()

	if err != nil {
		return nil, err
	}

	return &commentsRepo {
		log: repo.log,
		pc: repo,
	}, nil
}

func (cr commentsRepo) InsertComment(c *models.Comment) error {
	ctx, cancel := getContext()
	defer cancel()

	cr.log.Trace().Msg("beginning to insert comment")

	insertStr := `insert into "comments"("commentid", "description") values($1, $2)`

	_, err := cr.pc.pool.ExecContext(ctx, insertStr, c.CommentID, c.Description)
	if err != nil {
		cr.log.Error().
			Stack().
			Err(err).
			Msg("issue inserting into db")
		return err
	}

	cr.log.Debug().
		Str("commentId", c.CommentID).
		Str("description", c.Description).
		Msg("comment inserted")

	return nil
}

func (cr commentsRepo) GetAllComments() (*models.Comments, error) {
	cr.log.Trace().Msg("beginning to get all comments")

	comments := &models.Comments{
		Data: []*models.Comment{},
	}

	ctx, cancel := getContext()
	defer cancel()

	selectStr := `select * from comments`

	rows, err := cr.pc.pool.QueryContext(ctx, selectStr)
	if err != nil {
		cr.log.Error().
			Stack().
			Err(err).
			Msg("issue getting all rows of comments")
		return comments, err
	}

	var total int
	total = 0

	// loop through rows and scan to assign to column data to struct fields
	for rows.Next() {
		comment := &models.Comment{}
		if err := rows.Scan(&comment.CommentID, &comment.Description); err != nil {
			cr.log.Error().
				Stack().
				Err(err).
				Msg("row doesn't have the correct column data?")
			return comments, err
		}
		total = total + 1
		comments.Data = append(comments.Data, comment)
	}

	comments.Total = total

	cr.log.Debug().Msg("getting all comments successful")
	
	return comments, nil
}