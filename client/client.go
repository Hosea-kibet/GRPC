package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Hosea-kibet/GRPC_EXAMPLE/coffeshop_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to GRPC:", err)
	}
	defer conn.Close()

	// Create a client for the CoffeeShop service
	c := pb.NewCoffeeShopClient(conn)

	// Set a timeout for the context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the GetMenu RPC
	menuStream, err := c.GetMenu(ctx, &pb.MenuRequest{})
	if err != nil {
		log.Fatal("Error calling GetMenu function:", err)
	}

	done := make(chan bool)
	var items []*pb.Item

	// Goroutine to receive streamed menu items
	go func() {
		for {
			// Receive each item from the stream
			resp, err := menuStream.Recv()
			if err == io.EOF {
				// End of stream
				done <- true
				return
			}

			if err != nil {
				log.Fatalf("Cannot receive: %v", err)
			}

			// Assign the received items
			items = resp.Items
			log.Printf("Response received: %v", resp.Items) // Access Items correctly with uppercase 'I'
		}
	}()

	// Wait for the streaming to finish
	<-done

	// Place an order
	receipt, err := c.PlaceOrder(ctx, &pb.Order{Items: items})
	if err != nil {
		log.Fatalf("Error placing order: %v", err)
	}
	log.Printf("Order receipt: %v", receipt)

	// Check the status of the order
	status, err := c.OrderStatus(ctx, receipt)
	if err != nil {
		log.Fatalf("Error checking order status: %v", err)
	}
	log.Printf("Order status: %v", status)
}
