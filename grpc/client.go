package main

import (
	"context"
	"log"

	"github.com/khasyah-fr/golang-journey/grpc/chat"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello from Client"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %v", err)
	}

	log.Printf("Response from server: %v", response.Body)
}
