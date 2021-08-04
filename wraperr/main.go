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

// dao 层
func daoGetTodoById(id string) (out *Todo, err error) {
	// 根据id模拟一些错误
	switch id {
	case "0":
		err = sql.ErrNoRows
	case "1":
		err = sql.ErrConnDone
	default:
		out = &Todo{
			Id:    id,
			Title: "",
		}
	}

	if err != nil {
		// wrap err
		return nil, errors.Wrapf(err, "daoGetTodoByIdErr, id: %s", id)
	}

	return out, nil
}

// service 层
func serviceGetTodoById(id string) (*Todo, error) {
	todo, err := daoGetTodoById(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// handler 层
func handlerGetTodoById(ctx *gin.Context) {
	id := ctx.Param("id")

	todo, err := serviceGetTodoById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(404, gin.H{
				"code": 404,
				"msg":  "todo not found",
			})
		} else {
			// 其他错误需要记录日志
			fmt.Printf("%+v\n", err)
			ctx.JSON(500, gin.H{
				"code": 500,
				"msg":  "internal error",
			})
		}
		return
	}
	ctx.JSON(200, todo)
}

func main() {
	r := gin.New()

	r.GET("/todo/:id", handlerGetTodoById)

	_ = r.Run(":8000")
}
