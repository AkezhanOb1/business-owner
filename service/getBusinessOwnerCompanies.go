package service

import (
	"context"
	pb "github.com/AkezhanOb1/business-owner/api"
	db "github.com/AkezhanOb1/business-owner/repository"
)


//GetBusinessOwnerCompanies
func (s *Server) GetBusinessOwnerCompanies(ctx context.Context, request *pb.GetBusinessOwnerCompaniesRequest) (*pb.GetBusinessOwnerCompaniesResponse, error) {
	companies, err := db.GetBusinessOwnerCompaniesRepository(ctx, request.GetEmail())
	if err != nil {
		return nil, err
	}

	return companies, nil
}
