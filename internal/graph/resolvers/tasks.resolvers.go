package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"github.com/nicowernli/graphql-tutorial/internal/graph"
	graphmodels "github.com/nicowernli/graphql-tutorial/internal/graph/models"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input graphmodels.NewTask) (*graphmodels.Task, error) {
	panic(fmt.Errorf("not implemented: CreateTask - createTask"))
}

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, id string, name string, description *string) (*graphmodels.Task, error) {
	panic(fmt.Errorf("not implemented: UpdateTask - updateTask"))
}

// DeleteTask is the resolver for the deleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteTask - deleteTask"))
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context) ([]*graphmodels.Task, error) {
	panic(fmt.Errorf("not implemented: Tasks - tasks"))
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, id string) (*graphmodels.Task, error) {
	panic(fmt.Errorf("not implemented: Task - task"))
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }