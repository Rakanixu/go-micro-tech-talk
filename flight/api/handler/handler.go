package handler

import (
	"context"

	"github.com/Rakanixu/go-micro-tech-talk/flight/srv/proto/flight"
	restful "github.com/emicklei/go-restful"
)

type API struct {
	client proto_flight.ServiceClient
}

func NewAPI(client proto_flight.ServiceClient) *API {
	return &API{
		client: client,
	}
}

func (a *API) ReadAll(req *restful.Request, rsp *restful.Response) {
	r, err := a.client.Search(context.TODO(), &proto_flight.SearchRequest{
		Query: "{}",
	})
	if err != nil {
		rsp.WriteError(500, err)
		return
	}
	rsp.WriteEntity(r)
}

func (a *API) Read(req *restful.Request, rsp *restful.Response) {
	r, err := a.client.Read(context.TODO(), &proto_flight.ReadRequest{
		Id: req.PathParameter("id"),
	})

	if err != nil {
		rsp.WriteError(500, err)
		return
	}

	rsp.WriteEntity(r)
}
