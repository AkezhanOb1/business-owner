package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-owner/api"
	db "github.com/AkezhanOb1/business-owner/repository"
	"github.com/AkezhanOb1/business-owner/client/mq"
)


//CreateBusinessOwner is
func (s *Server) CreateBusinessOwner(ctx context.Context, request *pb.CreateBusinessOwnerRequest) (*pb.CreateBusinessOwnerResponse, error) {
	hashedPassword, err := hashPassword(request.GetBusinessOwnerPassword())
	if err != nil {
		return nil, err
	}
	request.BusinessOwnerPassword = hashedPassword

	businessOwner, err := db.CreateOwnerRepository(ctx, request)
	if err != nil {
		return nil, err
	}

	err = mq.BusinessRegistrationEmail(request.GetBusinessOwnerEmail())
	if err != nil {
		return nil, err
	}

	bindRequest := &pb.BindCompanyToBusinessOwnerRequest{
		BusinessOwnerID:      businessOwner.BusinessOwner.GetBusinessOwnerID(),
		BusinessCompanyID:    request.GetBusinessCompanyID(),
	}

	_, err = s.BindCompanyToBusinessOwner(context.Background(), bindRequest)
	if err != nil {
		return nil, err
	}

	return businessOwner, nil
}
