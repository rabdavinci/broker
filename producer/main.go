package main

import (
	"context"
	"log"

	"github.com/rabdavinci/broker/broker/proto"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := proto.NewMessageClient(conn)

	res, err := c.Message(context.Background(), &proto.Request{Topic: "test", Content: "Message"})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetResult())
}
