package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "github.com/danilkastar440/TRRP_third_task/proto"
)

const (
	host = "127.0.0.1"
	port = "50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)

	r, err := c.SquareRoot(context.Background(), &pb.Number{
		Value: 10,
	})
	if err != nil {
		log.Fatalf("Failed to save cars: %v", err)
	}
	log.Printf("Result: %f", r.GetValue())
}