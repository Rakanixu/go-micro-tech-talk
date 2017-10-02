package main

import (
	"log"

	"github.com/Rakanixu/go-micro-tech-talk/flight/srv/handler"
	"github.com/Rakanixu/go-micro-tech-talk/flight/srv/proto/flight"
	"github.com/Rakanixu/go-micro-tech-talk/lib/db"
	"github.com/Rakanixu/go-micro-tech-talk/lib/globals"
	_ "github.com/Rakanixu/go-micro-tech-talk/lib/plugins"
	"github.com/Rakanixu/go-micro-tech-talk/lib/wrappers"
)

func main() {
	// New service
	service := wrappers.NewService(globals.FLIGHT_SRV)

	if err := db.Init(globals.DB_URL); err != nil {
		log.Fatal(err)
	}

	// New service handler
	proto_flight.RegisterServiceHandler(service.Server(), new(handler.Service))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
