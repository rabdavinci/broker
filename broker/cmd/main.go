package main

import (
	"log"
	"net"
	"time"

	"github.com/rabdavinci/broker/broker/data/proto"
	"github.com/rabdavinci/broker/broker/service"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/cmd"
	"google.golang.org/grpc"
)

func main() {

	cmd.Init()
	// var cfg config.Config

	// if err := env.Parse(&cfg); err != nil {
	// 	log.Fatalf("parse config error with: %v", err)
	// }
	// if err := cfg.Validate(); err != nil {
	// 	log.Fatalf("validate config error with: %v", err)
	// }

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	s := grpc.NewServer()
	srv := &service.GRPCServer{}

	proto.RegisterMessageServer(s, srv)

	// cfg.GRPCAddress = "8080"
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("unable to create grpc listener: %v", err)
	}
	log.Println("domain grpc server was started on", ":8080")

	if err = s.Serve(l); err != nil {
		log.Fatalf("unable to start grpc server: %v", err)
	}

	<-time.After(time.Second * 20)
}
