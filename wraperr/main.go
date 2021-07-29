package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Todo struct {
	Id    string
	Title string
}

func daoGetTodoById(id string) (*Todo, error) {
	// 模拟 ErrNoRows
	err := sql.ErrNoRows
	// 在此处wrap err
	return nil, errors.Wrapf(err, "daoGetTodoByIdErr, id: %s", id)
}

func logicGetTodoById(id string) (*Todo, error) {
	todo, err := daoGetTodoById(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func handlerGetTodoById(ctx *gin.Context) {
	id := ctx.Param("id")

	todo, err := logicGetTodoById(id)
	if err != nil {
		// 在handler层处理err可以看到具体的错误堆栈
		fmt.Printf("%+v\n", err)
		ctx.AbortWithStatus(400)
		return
	}
	ctx.JSON(200, todo)
}

func main() {
	r := gin.New()

	r.GET("/todo/:id", handlerGetTodoById)

	_ = r.Run(":8000")
}
