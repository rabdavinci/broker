package main

import (
	"context"
	"log"

	"github.com/rabdavinci/broker/broker/data"
	"github.com/rabdavinci/broker/broker/data/proto"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := proto.NewMessageClient(conn)
	message := data.Message{
		Topic:   "qqq",
		Content: "qqq",
	}
	res, err := c.Message(context.Background(), message)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetResult())
}
