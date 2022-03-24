package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hellokvn/go-grpc-order-svc/pkg/client"
	"github.com/hellokvn/go-grpc-order-svc/pkg/db"
	"github.com/hellokvn/go-grpc-order-svc/pkg/pb"
	"github.com/hellokvn/go-grpc-order-svc/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	port := ":50053"
	h := db.Init("postgres://kevin@localhost:5432/order_svc")
	productSvc := client.InitProductServiceClient()

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Order Svc on", port)

	s := service.Server{
		H:          h,
		ProductSvc: productSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
