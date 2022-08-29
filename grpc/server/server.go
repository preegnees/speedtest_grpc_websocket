package main

import (
	"log"
	"net"
	"time"

	pb "grpc_vs_websocket/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type server struct {
	pb.MathServer
}

func main() {
	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMathServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}



func (s server) Max(stream pb.Math_MaxServer) error {
	
	mt, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		log.Println(mt["name"])
	}
	
	defer func() {
		_ = recover() 
	}()
	go func() {
		for {
			time.Sleep(1 * time.Second)
			_ = stream.Send(&pb.Response{Result: int32(time.Now().Second())})
		}
	}()
	for {
		select{
		case <- stream.Context().Done():
			stream.Context().Err()
		default:
			for {
				resp, _ := stream.Recv()
				log.Println(resp.Num)
			}
		}
	}
}