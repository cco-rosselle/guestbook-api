// copy/paste lol
package routes

import (
	"home/zellie/Code/guestbook-api/internal/models"
	"home/zellie/Code/guestbook-api/internal/settings"

	middleware "github.com/clearchanneloutdoor/token-middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	query "go.jtlabs.io/query"
)
 
func attachQueryOptions(ctx *gin.Context) {
	s := settings.Default()

	query, err := query.FromQuerystring(ctx.Request.URL.RawQuery)
	if err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}

	if query.Page == nil {
		query.Page = make(map[string]int)
	}

	if _, ok := query.Page["limit"]; !ok {
		query.Page["limit"] = s.Data.MaxLimit
	}

	if _, ok := query.Page["offset"]; !ok {
		query.Page["offset"] = 0
	}

	ctx.Set("QueryOptions", query)

	ctx.Next()
}

func errorReporter(ctx *gin.Context) {
	ctx.Next()
	detectedErrs := ctx.Errors.ByType(gin.ErrorTypeAny)

	if len(detectedErrs) == 0 {
		return
	}

	log.Debug().Int("errors", len(detectedErrs)).Msg("errors found")

	err := detectedErrs.Last().Err
	var parsedErr *models.APIError

	switch err := err.(type) {
		case *models.APIError:
			parsedErr = err

		case *middleware.AuthError:
			parsedErr = &models.APIError {
				Message: err.Message,
				Status: err.Status,
				Title: err.Title,
			}
		
		default:
			parsedErr = models.InternalServerError("")
	}

	ctx.JSON(parsedErr.Status, parsedErr)
	ctx.Abort()
}

func RegisterMiddleware(r *gin.Engine) {
	log.Debug().Msg("middleware registered for API")

	r.Use(attachQueryOptions)
	r.Use(errorReporter)
	r.Use(cors.New(cors.Config {
		AllowAllOrigins: true,
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"OPTIONS", "GET", "PATCH", "PUT", "POST", "DELETE"},
		ExposeHeaders: []string{"*", "Authorization"},
		AllowCredentials: true, // or nah?????
	}))
}