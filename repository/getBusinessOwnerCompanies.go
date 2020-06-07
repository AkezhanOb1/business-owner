package repository


import (
	"context"
	pb "github.com/AkezhanOb1/business-owner/api"
	"github.com/AkezhanOb1/business-owner/config"
	"github.com/jackc/pgx/v4"
)


//GetBusinessOwnerCompaniesRepository is a
func GetBusinessOwnerCompaniesRepository(ctx context.Context, email string) (*pb.GetBusinessOwnerCompaniesResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	sqlQuery := `SELECT business_company.id, business_company.name, business_company.category_id FROM business_company
					 INNER JOIN owners_businesses ob ON business_company.id = ob.business_company_id
					 INNER JOIN business_owner bo ON ob.business_owner_id = bo.id
				 WHERE bo.email=$1;`



	rows, err := conn.Query(ctx, sqlQuery, email)
	if err != nil {
		return nil, err
	}

	var companies pb.GetBusinessOwnerCompaniesResponse

	for rows.Next() {
		var company pb.BusinessCompany

		err = rows.Scan(
			&company.BusinessCompanyID,
			&company.BusinessCompanyName,
			&company.BusinessCompanyCategoryID,
		)

		if err != nil {
			return nil, err
		}

		companies.Companies = append(companies.Companies, &company)
	}

	return &companies, nil
}

