package main

import (
	"log"
	"net"
	"time"

	"github.com/rabdavinci/broker/broker/data"
	"github.com/rabdavinci/broker/broker/proto"
	"github.com/rabdavinci/broker/broker/service"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// GRPS SERVER SECTION START
	s := grpc.NewServer()
	srv := &service.GRPCServer{}

	// cfg.GRPCAddress = "8080"
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("unable to create grpc listener: %v", err)
	}
	log.Println("domain grpc server was started on", ":8080")

	if err = s.Serve(l); err != nil {
		log.Fatalf("unable to start grpc server: %v", err)
	}
	// GRPS SERVER SECTION END

	// RabbitMQ INIT START

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	brokerRepo := data.Broker(conn)

	domainService := data.NewService(brokerRepo)

	proto.RegisterMessageServer(s, srv, domainService)
	// RabbitMQ INIT END
	<-time.After(time.Second * 200)
}
