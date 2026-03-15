package graph

import (
	"context"
	"fmt"

	"github.com/keodevspace/go-graphql-api/graph/model"
)

// CreateTask is the resolver for the createTask mutation.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTaskInput) (*model.Task, error) {
	newTask := &model.Task{
		ID:          fmt.Sprintf("%d", len(r.Store.Tasks)+1),
		Title:       input.Title,
		IsCompleted: false,
	}

	r.Store.Tasks = append(r.Store.Tasks, newTask)
	return newTask, nil
}

// Tasks is the resolver for the tasks query.
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	return r.Store.Tasks, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)
