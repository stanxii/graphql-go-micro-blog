package main

import (
	"api/graph"
	"api/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/micro/go-log"
	"github.com/micro/go-micro/v2/client"

	_ "github.com/micro/go-plugins/registry/etcdv3/v2"

	"github.com/micro/go-micro/v2/web"
	articleProto "api/proto/article"
)


func main() {
	service := web.NewService(
		web.Name("go.micro.api.article"),
	)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// RPC client
	articleClient := articleProto.NewArticleSrvService("go.micro.srv.article", client.DefaultClient)

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{
			Resolvers:  &graph.Resolver{articleClient},
			Directives: generated.DirectiveRoot{},
			Complexity: generated.ComplexityRoot{},
		}),
	)

	//srv.Use(extension.FixedComplexityLimit(500))

	service.Handle("/", playground.Handler("GraphQL playground", "/query"))
	service.Handle("/query", srv)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
