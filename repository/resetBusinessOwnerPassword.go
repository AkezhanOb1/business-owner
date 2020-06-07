package repository


import (
	"context"
	pb "github.com/AkezhanOb1/business-owner/api"
	"github.com/AkezhanOb1/business-owner/config"

	"github.com/jackc/pgx/v4"
)


//ResetBusinessOwnerPassword is a repository that responsible to inserting data into the owner
//table in database
func ResetBusinessOwnerPassword(ctx context.Context, request *pb.ResetBusinessOwnerPasswordRequest) (*pb.ResetBusinessOwnerPasswordResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(context.Background())

	sqlQuery := `UPDATE business_owner SET
					password=$1 WHERE email=$2
				 RETURNING *;`

	row := conn.QueryRow(
		context.Background(),
		sqlQuery,
		request.BusinessOwnerPassword,
		request.BusinessOwnerEmail,
	)

	var password string
	var owner pb.BusinessOwner
	err = row.Scan(
		&owner.BusinessOwnerID,
		&owner.BusinessOwnerName,
		&owner.BusinessOwnerEmail,
		&password,
		&owner.BusinessOwnerPhoneNumberPrefix,
		&owner.BusinessOwnerPhoneNumber,
	)
	if err != nil {
		return nil, err
	}

	return &pb.ResetBusinessOwnerPasswordResponse{
		BusinessOwner: &owner,
	}, nil
}

