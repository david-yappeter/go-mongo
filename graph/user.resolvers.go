package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *userOpsResolver) Create(ctx context.Context, obj *model.UserOps, input model.NewUser) (*model.User, error) {
	return service.UserCreate(input)
}

func (r *userOpsResolver) CreateBatch(ctx context.Context, obj *model.UserOps, input []*model.NewUser) ([]*model.User, error) {
	return service.UserCreateBatch(input)
}

func (r *userOpsResolver) Update(ctx context.Context, obj *model.UserOps, input model.UpdateUser) (*model.User, error) {
	return service.UserUpdate(input)
}

func (r *userOpsResolver) Delete(ctx context.Context, obj *model.UserOps, id string) (string, error) {
	return service.UserDelete(id)
}

// UserOps returns generated.UserOpsResolver implementation.
func (r *Resolver) UserOps() generated.UserOpsResolver { return &userOpsResolver{r} }

type userOpsResolver struct{ *Resolver }
