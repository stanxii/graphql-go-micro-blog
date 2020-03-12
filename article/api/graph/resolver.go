//go:generate go run github.com/99designs/gqlgen
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
package graph

import articleProto "api/proto/article"

type Resolver struct{
	ArticleClient articleProto.ArticleSrvService
}

type ToInt64 struct {
	x int
}