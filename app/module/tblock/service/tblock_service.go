package service

import (
	"context"
	tblockv1 "lattice-manager-grpc/gen/tblock/v1"
	"time"
)

func NewTBlockServer() tblockv1.TBlockServiceServer {
	return &tBlockServer{}
}

type tBlockServer struct {
	tblockv1.UnimplementedTBlockServiceServer
}

func (t tBlockServer) Details(ctx context.Context, request *tblockv1.TBlockServiceDetailsRequest) (*tblockv1.TBlockServiceDetailsResponse, error) {
	return &tblockv1.TBlockServiceDetailsResponse{
		Hash:      request.Hash,
		Height:    1,
		Timestamp: uint64(time.Now().Unix()),
	}, nil
}
