package interfaces

import (
	"github.com/gin-gonic/gin"
)

type CommentsApi interface {
	TestApiFunc(*gin.Context)
	PostComment(*gin.Context)
	GetAll(*gin.Context)
}