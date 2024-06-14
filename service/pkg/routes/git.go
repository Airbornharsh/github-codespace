package routes

import (
	"github.com/airbornharsh/github-codespace/service/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func GitInit(r *gin.RouterGroup) {
	gitR := r.Group("/git")


	gitR.POST("/clone", handlers.GitClone)
	// gitR.DELETE("/delete", handlers.DeleteContainer)
}
