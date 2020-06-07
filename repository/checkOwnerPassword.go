package repository

import (
	"context"
	"github.com/AkezhanOb1/business-owner/config"
	"github.com/jackc/pgx/v4"
)


//GetOwnerPasswordRepository is a
func GetOwnerPasswordRepository(ctx context.Context, email string) (*string, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	sqlQuery := `SELECT password FROM business_owner WHERE email=$1;`

	var password string
	err = conn.QueryRow(ctx, sqlQuery, email).Scan(&password)
	if err != nil {
		return nil, err
	}

	return &password, nil
}

