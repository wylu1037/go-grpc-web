package service

import (
	"context"
	tblockv1 "lattice-manager-grpc/gen/tblock/v1"
)

func NewTBlockServer() tblockv1.TBlockServiceServer {
	return &tBlockServer{}
}

type tBlockServer struct {
	tblockv1.UnimplementedTBlockServiceServer
}

func (t tBlockServer) Details(ctx context.Context, request *tblockv1.TBlockServiceDetailsRequest) (*tblockv1.TBlockServiceDetailsResponse, error) {
	return nil, nil
}
