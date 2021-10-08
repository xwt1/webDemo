package routers

import (
	v1 "webDemo/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	user := v1.NewUser()
	program := v1.NewProgram()
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/users", user.Create)
		apiv1.DELETE("/users/:name", user.Delete)
		apiv1.PUT("/users/:name", user.Update)
		apiv1.GET("/users", user.Get)
		apiv1.GET("users/login", user.Login)

		apiv1.POST("/programs", program.Create)
		apiv1.DELETE("/programs/:program_id", program.Delete)
		apiv1.PUT("/programs/:program_id", program.Update)
		apiv1.GET("/programs/:program_id", program.Get)
		apiv1.GET("/programs", program.List)
	}
	return r
}
