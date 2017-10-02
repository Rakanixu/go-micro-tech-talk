package wrappers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Rakanixu/go-micro-tech-talk/lib/globals"
	"github.com/micro/go-micro"
	"github.com/micro/go-web"
)

func NewService(name string) micro.Service {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("oops can't get hostname")
	}

	md := map[string]string{
		"hostname": hostname,
	}

	sn := fmt.Sprintf("%s.srv.%s", globals.NAMESPACE, name)

	service := micro.NewService(
		micro.Name(sn),
		micro.Version("latest"),
		micro.Metadata(md),
		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(time.Second*30),
	)

	service.Init(
	/* 		micro.WrapClient(
	   			ContextClientWrapper(service),
	   		),
	   		micro.WrapSubscriber(
	   			NewContextSubscriberWrapper(service),
	   			NewAuthSubscriberWrapper(),
	   			NewAfterSubscriberWrapper(),
	   			NewQuotaSubscriberWrapper(sn),
	   			NewLogSubscriberWrapper(),
	   		),
	   		micro.WrapHandler(
	   			NewContextHandlerWrapper(service),
	   			NewAuthHandlerWrapper(),
	   			NewAfterHandlerWrapper(),
	   			NewQuotaHandlerWrapper(sn),
	   			NewLogHandlerWrapper(),
	   		), */
	)

	return service
}

func NewWebService(name string) web.Service {
	sn := fmt.Sprintf("%s.api.%s", globals.NAMESPACE, name)

	service := web.NewService(
		web.Name(sn),
	)

	service.Init()

	return service
}
