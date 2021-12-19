package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/alfieloakes/blog/log"
	"github.com/alfieloakes/blog/registry"
	"github.com/alfieloakes/blog/service"
)

func main() {
	log.Run("./app.log")

	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.LogService
	r.ServiceUrl = serviceAddress

	ctx, err := service.Start(context.Background(),

		host,
		port,
		r,
		log.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down log service")
}
