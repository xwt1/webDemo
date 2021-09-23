package routers

import ("github.com/gin-gonic/gin"
	"webDemo/internal/routers/api/v1"
)

func NewRouter() *gin.Engine{
	r := gin.New()
	r.Use((gin.Logger()))
	r.Use(gin.Recovery())

	user := v1.NewUser()
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/users", user.Create)
		apiv1.DELETE("/users/:name", user.Delete)
		apiv1.PUT("/users/:name", user.Update)
		apiv1.GET("/users", user.Get)
	}
	return r
}