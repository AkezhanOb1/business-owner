package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-owner/api"
	"github.com/AkezhanOb1/business-owner/client/mq"
)




//BindCompanyToBusinessOwner is
func (*Server) BusinessOwnerPasswordForgot(ctx context.Context, request *pb.BusinessOwnerPasswordForgotRequest) (*pb.BusinessOwnerPasswordForgotResponse, error) {
	err := mq.BusinessOwnerResetPassword(request.GetBusinessOwnerEmail())
	if err != nil {
		return nil, err
	}
	return &pb.BusinessOwnerPasswordForgotResponse{
		Success:  true,
	}, nil
}