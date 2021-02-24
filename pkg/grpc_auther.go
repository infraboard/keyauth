package pkg

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

var (
	// GRPCAuther todo
	GRPCAuther = &grpcAuther{}
)

// internal todo
type grpcAuther struct{}

func (a *grpcAuther) Filter(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println("filter:", info)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	return handler(ctx, req)
}
