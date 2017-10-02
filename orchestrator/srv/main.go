package main

import (
	"log"

	"github.com/Rakanixu/go-micro-tech-talk/lib/globals"
	_ "github.com/Rakanixu/go-micro-tech-talk/lib/plugins"
	"github.com/Rakanixu/go-micro-tech-talk/lib/wrappers"
	"github.com/Rakanixu/go-micro-tech-talk/orchestrator/srv/tasks"
)

func main() {
	// New service
	srv := wrappers.NewService(globals.ORCHESTRATOR_SRV)

	// Trigger flight indexing every 10 seconds
	tasks.ContinuosFlightIndexing(srv, 10)

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
