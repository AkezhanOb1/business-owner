package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-owner/api"
	db "github.com/AkezhanOb1/business-owner/repository"
)




//BindCompanyToBusinessOwner is
func (*Server) BindCompanyToBusinessOwner(ctx context.Context, request *pb.BindCompanyToBusinessOwnerRequest) (*pb.BindCompanyToBusinessOwnerResponse, error) {
	binding, err := db.BindCompanyToBusinessOwner(ctx, request)
	if err != nil {
		return nil, err
	}

	return binding, nil
}