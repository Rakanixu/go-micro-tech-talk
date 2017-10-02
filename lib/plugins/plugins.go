package plugins

import (
	// Plugins import
	_ "github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	_ "github.com/micro/go-plugins/transport/grpc"
	_ "github.com/micro/go-plugins/transport/tcp"

	// DB implementation
	_ "github.com/Rakanixu/go-micro-tech-talk/lib/db/elastic"
)
