package repos

import (
	"database/sql"
	"home/zellie/Code/guestbook-api/internal/models"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

type commentsRepo struct {
	log zerolog.Logger
	pc  *PostgresRepo
}

func NewCommentsRepo() (*commentsRepo, error) {
	repo, err := newPostgresRepo()

	if err != nil {
		return nil, err
	}

	return &commentsRepo{
		log: repo.log,
		pc:  repo,
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
		return nil, err
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
			return nil, err
		}
		total = total + 1
		comments.Data = append(comments.Data, comment)
	}

	comments.Total = total

	cr.log.Debug().Msg("getting all comments successful")

	return comments, nil
}

func (cr commentsRepo) DeleteComment(cid string) error {
	ctx, cancel := getContext()
	defer cancel()

	if found, err := cr.findComment(cid); !found {
		cr.log.Debug().Msg("error while searching commentId")
		return err
	}

	deleteStr := `delete from comments where commentid = $1`

	_, err := cr.pc.pool.ExecContext(ctx, deleteStr, cid)
	if err != nil {
		cr.log.Error().
			Stack().
			Err(err).
			Str("commentId", cid).
			Msg("issue deleting commentId")
		return err
	}

	cr.log.Debug().
		Str("commentId", cid).
		Msg("comment deleted")

	return nil
}

func (cr commentsRepo) findComment(cid string) (bool, error) {
	ctx, cancel := getContext()
	defer cancel()

	findStr := `select commentid, description from comments where commentid = $1`

	_, err := cr.pc.pool.QueryContext(ctx, findStr, cid)
	if err != nil {
		if err == sql.ErrNoRows {
			cr.log.Error().
				Stack().
				Err(err).
				Str("commentId", cid).
				Msg("cannot find commentId")
		}

		cr.log.Error().
			Stack().
			Err(err).
			Str("commentId", cid).
			Msg("issue finding commentId")

		return false, err
	}
	return true, nil
}
