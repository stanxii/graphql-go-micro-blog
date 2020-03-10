//go:generate go run github.com/99designs/gqlgen
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
package graph

import (
	userProto "api/proto/user"
)

type Resolver struct{
	UserClient userProto.UserSrvService
}

type ToInt64 struct {
	x int
}