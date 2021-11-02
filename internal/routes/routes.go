// package routes exposes function for registering
// routes to a supplied gin instance
package routes

import (
	v1 "home/zellie/Code/guestbook-api/internal/routes/v1"
	"home/zellie/Code/guestbook-api/internal/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Register(r *gin.Engine, cs interfaces.CommentsService) {
	RegisterMiddleware(r)
	log.With().Str("package", "routes").Logger()
	log.Debug().Msg("creating routing group for all /v1 resources")
	v1.Register(r, cs)
}