package repository

import (
	"context"

	"github.com/ca-nambara-teruaki/sample-app/ent"
)

type TaskRepositoryInterface interface {
	CreateTask(ctx context.Context, task *ent.Task) error
	GetTask(ctx context.Context, id int) (*ent.Task, error)
	ListTasks(ctx context.Context) ([]*ent.Task, error)
	UpdateTask(ctx context.Context, task *ent.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	client *ent.Client
}

func NewTaskRepository(client *ent.Client) TaskRepositoryInterface {
	return &taskRepository{
		client: client,
	}
}

func (r *taskRepository) CreateTask(ctx context.Context, task *ent.Task) error {
	return r.client.Task.Create().
		SetTitle(task.Title).
		SetDescription(task.Description).
		SetCreatedBy(task.CreatedBy).
		SetIsDeleted(task.IsDeleted).
		Exec(ctx)
}

func (r *taskRepository) GetTask(ctx context.Context, id int) (*ent.Task, error) {
	return r.client.Task.Get(ctx, id)
}

func (r *taskRepository) ListTasks(ctx context.Context) ([]*ent.Task, error) {
	return r.client.Task.Query().All(ctx)
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *ent.Task) error {
	return r.client.Task.UpdateOne(task).Exec(ctx)
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	return r.client.Task.DeleteOneID(id).Exec(ctx)
}
