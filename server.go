package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Hosea-kibet/GRPC_EXAMPLE/coffeshop_proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(menuRequest *pb.MenuRequest, srv pb.CoffeeShop_GetMenuServer) error {
	items := []*pb.Item{
		&pb.Item{
			Id:   "1",
			Name: "Cuppicino mix",
		},
		&pb.Item{
			Id:   "2",
			Name: "Cuppicino",
		},
		&pb.Item{
			Id:   "3",
			Name: "Black Cofeee",
		},
	}

	for i, _ := range items {
		srv.Send(&pb.Menu{
			Items: items[0 : i+1],
		})
	}
	return nil
}

func (s *server) PlaceOrder(context context.Context, order *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Id: "ABC1234",
	}, nil
}

func (s *server) OrderStatus(context context.Context, receipt *pb.Receipt) (*pb.OrderStatusResponse, error) {
	return &pb.OrderStatusResponse{
		OrderId: receipt.Id,
		Status:  "IN PROGRESS",
	}, nil
}

func main() {
	//Setup a listener on port 9001
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Failed to listen  :%v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCoffeeShopServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve :%s", err)
	}

}
