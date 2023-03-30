package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-list-cqrs-es/internal/ports/api"
)

func GetUser(app api.API) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId := c.Param("id")
		user, err := app.GetUser(c, userId)
		if err != nil {
			return
		}

		if err != nil {
			c.JSON(http.StatusNotFound, gin.Error{Err: err, Type: gin.ErrorTypePublic})
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

func CreateTodoList(app api.API) func(c *gin.Context) {
	return func(c *gin.Context) {
		var todolist api.CreateTodoListRequest
		err := c.BindJSON(&todolist)
		if err != nil {

			c.JSON(http.StatusBadRequest, gin.Error{Err: err, Type: gin.ErrorTypePublic})
			return
		}

		user, err := app.CreateTodoList(c, todolist)
		if err != nil {
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.Error{Err: err, Type: gin.ErrorTypePublic})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"user": user})
	}
}

func GetTodoList(app api.API) func(c *gin.Context) {
	return func(c *gin.Context) {

		todoListId := c.Param("id")

		todolist, err := app.GetTodoList(c, todoListId)
		if err != nil {
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.Error{Err: err, Type: gin.ErrorTypePublic})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"todolist": todolist})

	}
}
