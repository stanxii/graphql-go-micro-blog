package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/graph/generated"
	"api/graph/model"
	"context"
)

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	res, err := r.Query().User(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
