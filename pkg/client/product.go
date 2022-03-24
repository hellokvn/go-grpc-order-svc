package client

import (
	"context"
	"fmt"

	"github.com/hellokvn/go-grpc-order-svc/pkg/pb"
	"google.golang.org/grpc"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func InitProductServiceClient() ProductServiceClient {
	fmt.Println("Product Service Client")

	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := ProductServiceClient{
		Client: pb.NewProductServiceClient(cc),
	}

	return c
}

func (c *ProductServiceClient) FindOne(productId int32) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: productId,
	}

	return c.Client.FindOne(context.Background(), req)
}

func (c *ProductServiceClient) DecreaseStock(productId int32, orderId int32) (*pb.DecreaseStockResponse, error) {
	req := &pb.DecreaseStockRequest{
		Id:      productId,
		OrderId: orderId,
	}

	return c.Client.DecreaseStock(context.Background(), req)
}
