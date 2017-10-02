package main

import (
	"log"

	"github.com/micro/go-micro/server"

	"github.com/Rakanixu/go-micro-tech-talk/indexer/srv/subscriber"
	"github.com/Rakanixu/go-micro-tech-talk/lib/db"
	"github.com/Rakanixu/go-micro-tech-talk/lib/globals"
	_ "github.com/Rakanixu/go-micro-tech-talk/lib/plugins"
	"github.com/Rakanixu/go-micro-tech-talk/lib/wrappers"
)

func main() {
	// New service
	srv := wrappers.NewService(globals.INDEXER_SRV)

	if err := db.Init(globals.DB_URL); err != nil {
		log.Fatal(err)
	}

	if err := srv.Server().Subscribe(
		srv.Server().NewSubscriber(
			globals.INDEX_FLIGHT_TOPIC,
			subscriber.NewTaskHandler(100),
			server.SubscriberQueue(globals.INDEX_FLIGHT_QUEUE),
		),
	); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
