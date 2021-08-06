package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Todo struct {
	ID       int
	Title    string
	Detail   string
	Deadline time.Time
	Status   int8
}

type TodoRepo interface {
	ListTodo(ctx context.Context) ([]*Todo, error)
	CreateTodo(ctx context.Context, todo *Todo) error
	UpdateTodo(ctx context.Context, todo *Todo) error
}

type TodoUsecase struct {
	repo TodoRepo
	log  *log.Helper
}

func NewTodoUsecase(repo TodoRepo, logger log.Logger) *TodoUsecase {
	return &TodoUsecase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/todo")),
	}
}

func (uc *TodoUsecase) Create(ctx context.Context, todo *Todo) error {
	return uc.repo.CreateTodo(ctx, todo)
}

func (uc *TodoUsecase) List(ctx context.Context) ([]*Todo, error) {
	return uc.repo.ListTodo(ctx)
}

func (uc *TodoUsecase) Update(ctx context.Context, todo *Todo) error {
	return uc.repo.UpdateTodo(ctx, todo)
}
