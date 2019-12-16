package main

import (
	"gitTest2/handle"
	"gitTest2/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:9898")
	if err != nil {
		log.Fatal("net.Listen")
	}
	s := grpc.NewServer()
	proto.RegisterBookInfoServiceServer(s, &handle.BookInfoService{})
	log.Print("grpc server running...")
	if err := s.Serve(lis); err != nil {
		log.Fatal("s.Serve err", err)
	}
}
