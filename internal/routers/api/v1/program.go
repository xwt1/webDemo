package v1

import (
	"fmt"
	"webDemo/internal/service"
	"webDemo/pkg/app"
	"webDemo/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Program struct{}

func NewProgram() Program {
	return Program{}
}

func (program *Program) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
func (program *Program) List(c *gin.Context) {

	response := app.NewResponse(c)

	response.ToResponse(gin.H{})
	return
}

func (program *Program) Create(c *gin.Context) {
	param := service.CreateProgramRequest{}
	c.ShouldBind(&param)
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	err := svc.CreateProgram(&param)
	if err != nil {
		fmt.Printf("svc.CreateProgram err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateProgramFail)
		return
	}
	response.ToResponse(gin.H{"create": "ok"})
	return
}
func (program *Program) Update(c *gin.Context) {

}

func (program *Program) Delete(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
