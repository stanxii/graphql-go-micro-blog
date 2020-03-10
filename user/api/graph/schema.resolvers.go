package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/graph/generated"
	"api/graph/model"
	proto "api/proto/user"
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
)

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	fmt.Println("user:", id)
	_id, _ := graphql.UnmarshalInt64(id)
	res, err := r.UserClient.User(ctx, &proto.UserRequest{Id: _id})
	if err != nil {
		return nil, err
	}

	gender, _ := graphql.UnmarshalInt(res.Gender)

	return &model.User{
		ID:     id,
		Name:   &res.Name,
		Email:  &res.Emial,
		Gender: &gender,
		Avatar: &res.Avatar,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
