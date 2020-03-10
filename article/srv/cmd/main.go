package main

import (
	"srv/handler"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	_ "srv/pkg/mysql"
	article "srv/proto/article"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.article"),
		micro.Version("latest"),
		micro.Address(":50010"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	article.RegisterArticleSrvHandler(service.Server(), new(handler.ArticleSrv))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.article", service.Server(), new(subscriber.Article))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
