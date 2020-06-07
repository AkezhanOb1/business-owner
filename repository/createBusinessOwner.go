package repository


import (
	"context"
	pb "github.com/AkezhanOb1/business-owner/api"
	"github.com/AkezhanOb1/business-owner/config"
	"log"
	"strings"

	"github.com/jackc/pgx/v4"
)


//CreateOwnerRepository is a repository that responsible to inserting data into the owner
//table in database
func CreateOwnerRepository(ctx context.Context, request *pb.CreateBusinessOwnerRequest) (*pb.CreateBusinessOwnerResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer func() {
		err =  conn.Close(context.Background())
		if err != nil {
			log.Println(err)
		}
	}()

	sqlQuery := `INSERT INTO business_owner (
						name, email, password, phone_number_prefix, phone_number)
				 VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	var ownerID int64
	ownerName := request.GetBusinessOwnerName()
	ownerEmail := strings.ToLower(request.GetBusinessOwnerEmail())
	ownerPassword := request.GetBusinessOwnerPassword()
	ownerPhoneNumberPrefix := request.GetBusinessOwnerPhoneNumberPrefix()
	ownerPhoneNumber := request.GetBusinessOwnerPhoneNumber()

	row := conn.QueryRow(
		context.Background(),
		sqlQuery,
		ownerName,
		ownerEmail,
		ownerPassword,
		ownerPhoneNumberPrefix,
		ownerPhoneNumber,
	)

	err = row.Scan(&ownerID)
	if err != nil {
		return nil, err
	}

	return &pb.CreateBusinessOwnerResponse{
		BusinessOwner: &pb.BusinessOwner{
			BusinessOwnerID:                ownerID,
			BusinessOwnerName:              ownerName,
			BusinessOwnerEmail:             ownerEmail,
			BusinessOwnerPhoneNumberPrefix: ownerPhoneNumberPrefix,
			BusinessOwnerPhoneNumber:       ownerPhoneNumber,
		},
	}, nil
}

