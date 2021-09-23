package v1

import "github.com/gin-gonic/gin"
type User struct{}

func NewUser() User{
	return User{}
}

func (user *User)Get(c *gin.Context){}
func (user *User)Delete(c *gin.Context){}
func (user *User)Create(c *gin.Context){}
func (user *User)Update(c *gin.Context){}