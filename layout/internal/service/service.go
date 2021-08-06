package service

import (
	v1 "github.com/dch1228/gobestpractices/layout/api/todo/v1"
	"github.com/dch1228/gobestpractices/layout/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewTodoService)

type TodoService struct {
	v1.UnimplementedTodoServiceServer

	todo *biz.TodoUsecase
	log  *log.Helper
}
