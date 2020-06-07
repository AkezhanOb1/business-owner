package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-owner/api"
	db "github.com/AkezhanOb1/business-owner/repository"
)

//CheckOwnerPassword is
func (*Server) CheckOwnerPassword(ctx context.Context, request *pb.CheckOwnerPasswordRequest) (*pb.CheckOwnerPasswordResponse, error) {
	password, err := db.GetOwnerPasswordRepository(ctx, request.GetEmail())
	if err != nil {
		return nil, err
	}

	err = comparePassword(*password, request.GetPassword())
	if err != nil {
		return nil, err
	}
	return &pb.CheckOwnerPasswordResponse{
		Valid:true,
	}, nil
}