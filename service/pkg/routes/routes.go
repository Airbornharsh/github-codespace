package routes

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	api := r.Group("/api")

	GitInit(api)
	CodeInit(api)
}
