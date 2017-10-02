package main

import (
	"fmt"
	"log"

	"github.com/Rakanixu/go-micro-tech-talk/flight/api/handler"
	"github.com/Rakanixu/go-micro-tech-talk/flight/srv/proto/flight"
	"github.com/Rakanixu/go-micro-tech-talk/lib/globals"
	"github.com/Rakanixu/go-micro-tech-talk/lib/wrappers"
	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
)

func main() {
	srv := wrappers.NewWebService(globals.FLIGHT_SRV)
	api := handler.NewAPI(proto_flight.NewServiceClient(
		"com.go-micro-tech-talk.srv.flight",
		client.NewClient(),
	))
	ws := new(restful.WebService)
	wc := restful.NewContainer()
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path(fmt.Sprintf("/%s", globals.FLIGHT_SRV))
	ws.Route(ws.GET("/").To(api.ReadAll))
	ws.Route(ws.GET("/{id}").To(api.Read))
	wc.Add(ws)

	srv.Handle("/", wc)

	// Run server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
