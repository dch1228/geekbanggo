package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Todo struct {
	Title string
}

func daoGetTodo() (*Todo, error) {
	return nil, sql.ErrNoRows
}

func logicSetTitle() (*Todo, error) {
	todo, err := daoGetTodo()
	if err != nil {
		return nil, err
	}

	todo.Title = "title"
	return todo, nil
}

func handlerGetTodo(ctx *gin.Context) {
	todo, err := logicSetTitle()
	if err != nil {
		fmt.Printf("%+v\n", errors.Wrap(err, "get todo error"))
		ctx.AbortWithStatus(400)
		return
	}
	ctx.JSON(200, todo)
}

func main() {
	r := gin.New()

	r.GET("/todo", handlerGetTodo)

	_ = r.Run(":8000")
}
