package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-owner/api"
	db "github.com/AkezhanOb1/business-owner/repository"
)




//BindCompanyToBusinessOwner is
func (*Server) ResetBusinessOwnerPassword(ctx context.Context, request *pb.ResetBusinessOwnerPasswordRequest) (*pb.ResetBusinessOwnerPasswordResponse, error) {
	hashedPassword, err := hashPassword(request.GetBusinessOwnerPassword())
	if err != nil {
		return nil, err
	}
	request.BusinessOwnerPassword = hashedPassword

	owner, err := db.ResetBusinessOwnerPassword(ctx, request)
	if err != nil {
		return nil, err
	}

	return owner, nil
}