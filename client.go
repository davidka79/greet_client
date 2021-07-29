package main

import (
	"fmt"
	"github.com/davidka79/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("hello I2m a client")

	cc, err := grpc.Dial("localhost:500551", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("its an error %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGteetServiceClient(cc)

	fmt.Printf("%f", c)
}
