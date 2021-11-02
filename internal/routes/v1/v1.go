// route registers go here
package v1

import (
	"home/zellie/Code/guestbook-api/internal/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Register(r *gin.Engine, cs interfaces.CommentsService) {
	v1 := r.Group("/v1")

	commentsApi := NewCommentsApi(cs)

	log.Debug().Str("package", "v1").Msg("registering /v1 resources")

	v1.GET("/test", commentsApi.TestApiFunc)
	v1.POST("/comments", commentsApi.PostComment)
	v1.GET("/comments", commentsApi.GetAll)
}