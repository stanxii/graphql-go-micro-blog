package main

import (
	"srv/handler"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
	//"github.com/micro/go-micro/v2/registry"
	//"github.com/micro/go-micro/v2/registry/etcd"
	//"github.com/micro/go-micro/v2/config/cmd"

	_ "srv/pkg/mysql"
	article "srv/proto/article"
)

//func NewRegistry(opts ...registry.Option) registry.Registry {
//	return etcd.NewRegistry(opts...)
//}

func main() {
	//cmd.DefaultRegistries["etcdv3"] = etcd.NewRegistry
	//reg := NewRegistry(func(options *registry.Options) {
	//	options.Addrs = []string{
	//		"http://127.0.0.1:2380",
	//	}
	//})

	// New Service
	service := micro.NewService(
		//micro.Registry(reg),
		micro.Name("go.micro.srv.article"),
		micro.Version("latest"),
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
