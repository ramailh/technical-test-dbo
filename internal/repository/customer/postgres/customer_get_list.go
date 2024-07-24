package postgres

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

var queryGetListCustomer = `
	SELECT customer_id, first_name, last_name,  email, phone, created_at, updated_at
	FROM customer
`

var customerOrderFields = map[string]bool{
	"customer_id": true, "first_name": true, "last_name": true, "email": true, "phone": true, "created_at": true,
}

var orderMap = map[string]bool{
	"asc":  true,
	"desc": true,
}

func (repo *customerRepository) GetList(req *model.PaginationRequest) ([]*model.Customer, *model.Meta, error) {
	var sqlWhere, sqlOrderBy string
	var values []interface{}
	if req.Search != "" {
		values = append(values, fmt.Sprintf("%%%s%%", req.Search))
		sqlWhere = "WHERE "
		sqlWhere += "first_name LIKE $1 "
		sqlWhere += "OR last_name LIKE $1 "
		sqlWhere += "OR email LIKE $1 "
		sqlWhere += "OR phone LIKE $1 "
	}

	if req.OrderBy != "" && req.Order != "" {
		if orderMap[strings.ToLower(req.Order)] && customerOrderFields[req.OrderBy] {
			sqlOrderBy = fmt.Sprintf(" ORDER BY %s %s ", req.OrderBy, req.Order)
		}
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	if req.Limit > 100 {
		req.Limit = 100
	}

	if req.Page == 0 {
		req.Page = 1
	}

	offset := (req.Page - 1) * req.Limit
	sqlLimitOffset := fmt.Sprintf(" LIMIT %d OFFSET %d ", req.Limit, offset)

	query := queryGetListCustomer + sqlWhere + sqlOrderBy + sqlLimitOffset
	res, err := repo.db.Query(query, values...)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	defer res.Close()

	var results []*model.Customer
	for res.Next() {
		var result model.Customer
		err := res.Scan(
			&result.CustomerID, &result.FirstName, &result.LastName,
			&result.Email, &result.Phone, &result.CreatedAt, &result.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return nil, nil, err
		}
	}

	total, err := repo.countCustomer(sqlWhere)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	meta := &model.Meta{
		Total:       total,
		Limit:       req.Limit,
		CurrentPage: req.Page,
		TotalPages:  int(math.Ceil(float64(total) / float64(req.Limit))),
	}

	return results, meta, nil
}

func (repo *customerRepository) countCustomer(sqlWhere string) (int, error) {
	var count int
	err := repo.db.QueryRow("SELECT COUNT(*) FROM customer " + sqlWhere).Scan(&count)
	if err != nil {
		log.Println(err)
		return count, err
	}

	return count, nil
}
