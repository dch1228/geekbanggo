package data

import (
	"context"

	"github.com/dch1228/gobestpractices/layout/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

var _ biz.TodoRepo = (*todoRepo)(nil)

type todoRepo struct {
	data *Data
	log  *log.Helper
}

func NewTodoRepo(data *Data, logger log.Logger) biz.TodoRepo {
	return &todoRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/todo")),
	}
}

func (r *todoRepo) ListTodo(ctx context.Context) ([]*biz.Todo, error) {
	ps, err := r.data.db.Todo.Query().All(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	out := make([]*biz.Todo, 0, len(ps))
	for _, p := range ps {
		out = append(out, &biz.Todo{
			ID:       p.ID,
			Title:    p.Title,
			Detail:   p.Detail,
			Deadline: p.Deadline,
			Status:   p.Status,
		})
	}
	return out, nil
}

func (r *todoRepo) CreateTodo(ctx context.Context, todo *biz.Todo) error {
	_, err := r.data.db.Todo.
		Create().
		SetTitle(todo.Title).
		SetDetail(todo.Detail).
		SetDeadline(todo.Deadline).
		SetStatus(todo.Status).
		Save(ctx)
	return errors.WithStack(err)
}

func (r *todoRepo) UpdateTodo(ctx context.Context, todo *biz.Todo) error {
	p, err := r.data.db.Todo.Get(ctx, todo.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = p.Update().
		SetTitle(todo.Title).
		SetDetail(todo.Detail).
		SetDeadline(todo.Deadline).
		SetStatus(todo.Status).
		Save(ctx)
	return errors.WithStack(err)
}
