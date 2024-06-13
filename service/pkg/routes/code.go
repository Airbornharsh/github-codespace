package routes

import (
	"github.com/airbornharsh/github-codespace/service/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func CodeInit(r *gin.RouterGroup) {

	r.PUT("/code", handlers.CodeUpdate)
}
