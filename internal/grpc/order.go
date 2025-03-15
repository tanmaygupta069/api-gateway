package Grpc

import (
    OrderPb "github.com/tanmaygupta069/order-service-go/generated/order"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func NewOrderClient(addr string) (OrderPb.OrderServiceClient, error) {
    conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        return nil, err
    }
    return OrderPb.NewOrderServiceClient(conn), nil
}