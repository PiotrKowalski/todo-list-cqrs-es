package rest

import (
	"github.com/gin-gonic/gin"
	"todo-list-cqrs-es/internal/ports/api"
)

type Adapter struct {
	r *gin.Engine
}

func NewRESTService(application api.API) Adapter {
	r := gin.Default()

	v1 := r.Group("/v1")
	v1.GET("/user/:id", GetUser(application))
	v1.POST("/todolist", CreateTodoList(application))
	v1.GET("/todolist/:id", GetTodoList(application))

	return Adapter{r: r}
}

func (a *Adapter) Run() error {

	err := a.r.Run("localhost:9000")
	if err != nil {
		return err
	}

	return nil
}
