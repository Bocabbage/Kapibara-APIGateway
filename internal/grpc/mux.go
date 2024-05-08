package grpc

import (
	"context"
	mikananipb "kapibara-apigateway/internal/grpc/mikanani"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateMikananiMux(grpcSvcAddr string) (*runtime.ServeMux, error) {

	conn, err := grpc.NewClient(
		grpcSvcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	gwmux := runtime.NewServeMux()
	err = mikananipb.RegisterMikananiServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		return nil, err
	}

	return gwmux, nil
}
