package main

import (
	"context"
	"log"
	"time"

	pb "grpc_vs_websocket/grpc/proto"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	client := pb.NewMathClient(conn)
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("name", "radmir"),
	)
	stream, err := client.Max(ctx) // Не нудно туда ничего отправлять вообще, кроме контекста
	if err != nil {
		log.Println(err)
	}

	go func() {
		req := pb.Request{Num: 2}
		err := stream.Send(&req)
		if err != nil {
			// проверить на EOF
			log.Println(err)
		}
	}()

	done := make(chan struct{})
	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				// проверить на EOF
				log.Println(err)
				return
			}
			log.Println(resp.Result)
		}
	}()

	select{
	case <- time.After(3 * time.Second):
		panic("aaa")
	case <-done:
		log.Println("hahahha")
	}
}
