package Grpc

import (
    auth_pb "github.com/tanmaygupta069/auth-service-go/generated"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func NewAuthClient(addr string) (auth_pb.AuthServiceClient, error) {
    conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        return nil, err
    }
    return auth_pb.NewAuthServiceClient(conn), nil
}