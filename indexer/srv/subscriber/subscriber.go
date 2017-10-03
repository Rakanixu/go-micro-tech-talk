package subscriber

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Rakanixu/go-micro-tech-talk/lib/db"
	"github.com/Rakanixu/go-micro-tech-talk/lib/globals"
	messages "github.com/Rakanixu/go-micro-tech-talk/lib/messages/indexing"
	"github.com/Rakanixu/go-micro-tech-talk/lib/model/flight"
	"golang.org/x/net/context"
)

func NewTaskHandler(workers int) *taskHandler {
	var i initTaskHandler
	i = startWorkers

	t := &taskHandler{
		msgChan: make(chan msgChan, 1000000),
		workers: workers,
	}

	i(t)

	return t
}

type taskHandler struct {
	msgChan chan msgChan
	workers int
}

type msgChan struct {
	msg *messages.IndexFlightsMessage
	ctx context.Context
	err chan error
}

type initTaskHandler func(t *taskHandler)

func (e *taskHandler) IndexFlights(ctx context.Context, msg *messages.IndexFlightsMessage) error {
	c := msgChan{
		msg: msg,
		ctx: ctx,
		err: make(chan error),
	}
	// Queue internally
	e.msgChan <- c

	return <-c.err
}

func (e *taskHandler) queueListener(wID int) {
	for m := range e.msgChan {
		if err := processmsg(m); err != nil {
			log.Println("Error processing index flight message: ", err)
			m.err <- err
		}
		// Successful
		m.err <- nil
	}
}

func startWorkers(t *taskHandler) {
	for i := 0; i < t.workers; i++ {
		go t.queueListener(i)
	}
}

func processmsg(m msgChan) error {
	log.Println("Indexer processing msg. Data Origin: ", m.msg.Origin)

	req, err := http.NewRequest(http.MethodGet, m.msg.Origin, nil)
	if err != nil {
		return err
	}

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	var flights []flight.Flight
	if err := json.NewDecoder(rsp.Body).Decode(&flights); err != nil {
		return err
	}

	for _, v := range flights {
		b, err := json.Marshal(&v)
		if err != nil {
			return err
		}

		if err := db.Index(globals.INDEX_FLIGT, globals.TYPE_FLIGHT, v.GUID, string(b)); err != nil {
			return err
		}

		log.Println("Flight indexed with ID ", v.GUID)
	}

	return nil
}
