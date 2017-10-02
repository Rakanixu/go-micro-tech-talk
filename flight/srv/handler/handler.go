package handler

import (
	"encoding/json"
	"errors"

	proto_flight "github.com/Rakanixu/go-micro-tech-talk/flight/srv/proto/flight"
	"github.com/Rakanixu/go-micro-tech-talk/lib/db"
	"github.com/Rakanixu/go-micro-tech-talk/lib/globals"
	"github.com/Rakanixu/go-micro-tech-talk/lib/model/flight"
	"golang.org/x/net/context"
)

// Service ..
type Service struct{}

// Read ...
func (s *Service) Read(ctx context.Context, req *proto_flight.ReadRequest, rsp *proto_flight.ReadResponse) error {
	data, err := db.Read(globals.INDEX_FLIGT, globals.TYPE_FLIGHT, req.Id)
	if err != nil {
		return err
	}

	var f *flight.Flight
	raw, ok := data.(*json.RawMessage)
	if !ok {
		return errors.New("Error getting decoding data")
	}

	if err := json.Unmarshal(*raw, &f); err != nil {
		return err
	}

	rsp.Flight = &proto_flight.Flight{
		Guid:     f.GUID,
		Origin:   f.Origin,
		Destiny:  f.Destiny,
		Aircraft: f.Aircraft,
	}

	return nil
}

// Search ...
func (s *Service) Search(ctx context.Context, req *proto_flight.SearchRequest, rsp *proto_flight.SearchResponse) error {
	data, err := db.Search(globals.INDEX_FLIGT, globals.TYPE_FLIGHT, req.Query)
	if err != nil {
		return err
	}

	for _, v := range data {
		var f *flight.Flight
		raw, ok := v.(*json.RawMessage)
		if !ok {
			return errors.New("Error getting decoding data")
		}

		if err := json.Unmarshal(*raw, &f); err != nil {
			return err
		}

		rsp.Flights = append(rsp.Flights, &proto_flight.Flight{
			Guid:     f.GUID,
			Origin:   f.Origin,
			Destiny:  f.Destiny,
			Aircraft: f.Aircraft,
		})
	}

	return nil
}

// Health ...
func (s *Service) Health(ctx context.Context, req *proto_flight.HealthRequest, rsp *proto_flight.HealthResponse) error {
	rsp.Info = "OK"

	return nil
}
