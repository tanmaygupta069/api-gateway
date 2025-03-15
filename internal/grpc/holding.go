package Grpc

import (
    HoldingPb "github.com/tanmaygupta069/order-service-go/generated/holding"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func NewHoldingClient(addr string) (HoldingPb.HoldingServiceClient, error) {
    conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        return nil, err
    }
    return HoldingPb.NewHoldingServiceClient(conn), nil
}