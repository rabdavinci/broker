package service

import (
	"context"
	"log"

	"github.com/rabdavinci/broker/broker/data/proto"
)

type GRPCServer struct {
	proto.UnimplementedMessageServer
}

func (s *GRPCServer) Message(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	log.Println(req.Topic)
	return &proto.Response{Result: false}, nil
}
