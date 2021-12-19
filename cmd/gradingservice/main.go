package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/alfieloakes/blog/grades"
	"github.com/alfieloakes/blog/log"
	"github.com/alfieloakes/blog/registry"
	"github.com/alfieloakes/blog/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.GradingService
	r.ServiceURL = serviceAddress
	r.RequiredServices = []registry.ServiceName{registry.LogService}
	r.ServiceUpdateURL = r.ServiceURL + "/services"
	r.HeartbeatURL = r.ServiceURL + "/heartbeat"

	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at %v\n", logProvider)
		log.SetClientLogger(logProvider, registry.GradingService)
	}
	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
