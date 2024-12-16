package main

import (
	"log"
	pb "github.com/praveent04/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))  // using a insecure connection here

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names:= &pb.NamesList{
		Names: []string{"Praveen","Akhil","Alice","Bob"},
	}
	log.Println("1st one")
	callSayHello(client) // unary api
	log.Println("2nd one")
	callSayHelloServerStream(client, names)
	log.Println("3rd one")
	callSayHelloClientStream(client, names)
	log.Println("4th one")
	callSayHelloBidirectionalStream(client, names)
	
}