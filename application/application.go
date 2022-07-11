package application

import (
	"context"
	"sut-notification-go/lib/pkg/db"

	"google.golang.org/grpc"
)

type Application struct {
	DbClients   db.Handler
	GrpcServer  *grpc.Server
	GrpcClients map[string]*grpc.ClientConn
	Context     context.Context
}
