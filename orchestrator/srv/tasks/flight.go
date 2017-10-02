package tasks

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Rakanixu/go-micro-tech-talk/lib/globals"
	messages "github.com/Rakanixu/go-micro-tech-talk/lib/messages/indexing"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

// ContinuosFlightIndexing ..
func ContinuosFlightIndexing(srv micro.Service, seconds time.Duration) {
	rand.Seed(time.Now().UTC().UnixNano())

	ticker := time.NewTicker(seconds * time.Second)
	q := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("Orchestrator Sending Message on ", globals.INDEX_FLIGHT_TOPIC)
				if err := srv.Client().Publish(context.Background(), srv.Client().NewPublication(
					globals.INDEX_FLIGHT_TOPIC,
					&messages.IndexFlightsMessage{
						Handler: globals.DEFAULT,
						Origin:  globals.FLIGTHS_DATA_ORIGIN + "flights_data_" + strconv.Itoa(rand.Intn(2)+1) + ".json",
					},
				)); err != nil {
					log.Println(err)
				}

				log.Println("Orchestrator Sending Message on ", globals.DEFAULT)
				if err := srv.Client().Publish(context.Background(), srv.Client().NewPublication(
					globals.DEFAULT,
					&messages.IndexFlightsMessage{
						Handler: globals.DEFAULT,
						Origin:  globals.DEFAULT,
					},
				)); err != nil {
					log.Println(err)
				}
			case <-q:
				ticker.Stop()
				return
			}
		}
	}()
}
