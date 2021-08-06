package service

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	v1 "github.com/dch1228/gobestpractices/layout/api/todo/v1"

	"github.com/dch1228/gobestpractices/layout/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

func NewTodoService(todo *biz.TodoUsecase, logger log.Logger) *TodoService {
	return &TodoService{
		todo: todo,
		log:  log.NewHelper(log.With(logger, "module", "service/todo")),
	}
}

func (s *TodoService) CreateTodo(ctx context.Context, in *v1.CreateTodoReq) (*v1.CreateTodoReply, error) {
	p := in.GetTodo()
	_ = s.todo.Create(ctx, &biz.Todo{
		Title:    p.GetTitle(),
		Detail:   p.GetDetail(),
		Deadline: p.GetDeadline().AsTime(),
		Status:   int8(p.GetStatus()),
	})
	return &v1.CreateTodoReply{}, nil
}
func (s *TodoService) UpdateTodo(ctx context.Context, in *v1.UpdateTodoReq) (*v1.UpdateTodoReply, error) {
	p := in.GetTodo()
	err := s.todo.Update(ctx, &biz.Todo{
		ID:       int(p.GetId()),
		Title:    p.GetTitle(),
		Detail:   p.GetDetail(),
		Deadline: p.GetDeadline().AsTime(),
		Status:   int8(p.GetStatus()),
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateTodoReply{}, nil
}

func (s *TodoService) ListTodo(ctx context.Context, _ *v1.ListTodoReq) (*v1.ListTodoReply, error) {
	ps, err := s.todo.List(ctx)
	if err != nil {
		return nil, err
	}
	reply := &v1.ListTodoReply{}
	for _, p := range ps {
		reply.Results = append(reply.Results, &v1.Todo{
			Id:       int64(p.ID),
			Title:    p.Title,
			Detail:   p.Detail,
			Deadline: timestamppb.New(p.Deadline),
			Status:   int32(p.Status),
		})
	}
	return reply, nil
}
