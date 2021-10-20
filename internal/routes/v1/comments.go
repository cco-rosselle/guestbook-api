package v1

import (
	"net/http"

	"home/zellie/Code/guestbook-api/internal/interfaces"
	"home/zellie/Code/guestbook-api/internal/models"
	"home/zellie/Code/guestbook-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type commentsApi struct {
	log zerolog.Logger
	commentsService *services.CommentsService
}

// returns a new controller CommentsApi for interacting with CommentsService
func NewCommentsApi(cs *services.CommentsService) interfaces.CommentsApi {
	return &commentsApi {
		log: log.With().
			Str("package", "v1").
			Str("component", "comments api controller").
			Logger(),
		commentsService: cs,
	}
}

// the (a commentsApi) is a receiver
// like oop you need to call the instance of a class
// to use its methods
// mirrored getproductsbyid in capex-api
func (ca commentsApi) TestApiFunc(ctx *gin.Context) {
	ca.log.Trace().Msg("test comment api function was reached")

	err := ca.commentsService.TestServiceFunc()
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, ca)
}

func (ca commentsApi) PostComment(ctx *gin.Context) {
	ca.log.Debug().Msg("post comment request received")

	body := &models.Comment{}
	if err := ctx.ShouldBindJSON(body); err != nil {
		ctx.Error(models.BadRequestError("request body is not proper json"))
		ctx.AbortWithStatus(400)
		return
	}

	if err := ca.commentsService.PostComment(body); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return	
	}

	ctx.JSON(http.StatusCreated, body)
}


func (ca commentsApi) GetAll(ctx *gin.Context) {
	ca.log.Debug().Msg("get all comments request received")
	
	comments, err := ca.commentsService.GetAllComments()
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	} 

	ctx.JSON(http.StatusOK, comments)
}