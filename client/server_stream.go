package main

import (
	"context"
	"io"
	"log"

	pb "github.com/praveent04/grpc-demo/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList){
	log.Printf("streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(),names)
	
	if err != nil{
		log.Fatalf("could no tsend names: %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatalf("error while streaming %v", err)
		}
		log.Println(message)
	}
	log.Printf("streaming finished")
}