package repository

import (
	"context"
	pb "github.com/AkezhanOb1/business-owner/api"
	"github.com/AkezhanOb1/business-owner/config"
	"github.com/jackc/pgx/v4"
)


//BindCompanyToBusinessOwner is a repository that responsible to inserting data into the owner table in database
func BindCompanyToBusinessOwner(ctx context.Context, request *pb.BindCompanyToBusinessOwnerRequest) (*pb.BindCompanyToBusinessOwnerResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	sqlQuery := `INSERT INTO owners_businesses(
						business_owner_id, business_company_id)
                 VALUES ($1, $2) RETURNING true;`

	row := conn.QueryRow(ctx, sqlQuery,request.GetBusinessOwnerID(), request.GetBusinessCompanyID())

	var success bool
	err = row.Scan(&success)
	if err != nil {
		return nil, err
	}

	return &pb.BindCompanyToBusinessOwnerResponse{
		BindSuccess: success,
	}, nil
}

