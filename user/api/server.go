package main

import (
	//"github.com/99designs/gqlgen/graphql/handler/extension"
	"api/graph"
	"api/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/micro/go-log"
	"github.com/micro/go-micro/v2/client"

	"github.com/micro/go-micro/v2/web"
	userProto "api/proto/user"
)

func main() {
	service := web.NewService(
		web.Name("go.micro.api.greeter"),
		web.Version("latest"),
		web.Address(":8085"),
	)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// RPC client
	userClient := userProto.NewUserSrvService("go.micro.srv.user", client.DefaultClient)

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{UserClient: userClient},
			},
		),
	)

	//srv.Use(extension.FixedComplexityLimit(500))

	service.Handle("/", playground.Handler("GraphQL playground", "/query"))
	service.Handle("/query", srv)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)

	}
}
