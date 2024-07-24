package postgres

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

var queryGetListOrder = `
	SELECT order_id, customer_id, order_date, status, created_at, updated_at
	FROM order
`

var orderFields = map[string]bool{
	"order_id": true, "customer_id": true, "item": true, "order_date": true, "status": true, "created_at": true,
}

var orderMap = map[string]bool{
	"asc":  true,
	"desc": true,
}

func (repo *orderRepository) GetList(req *model.PaginationRequest) ([]*model.Order, *model.Meta, error) {
	var sqlWhere, sqlOrderBy string
	var values []interface{}
	if req.Search != "" {
		values = append(values, fmt.Sprintf("%%%s%%", req.Search))
		sqlWhere = "WHERE "
		sqlWhere += "item LIKE $1 "
	}

	if req.OrderBy != "" && req.Order != "" {
		if orderMap[strings.ToLower(req.Order)] && orderFields[req.OrderBy] {
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

	query := queryGetListOrder + sqlWhere + sqlOrderBy + sqlLimitOffset
	res, err := repo.db.Query(query, values...)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	defer res.Close()

	var results []*model.Order
	for res.Next() {
		var result model.Order
		err := res.Scan(
			&result.OrderID, &result.CustomerID, &result.OrderDate, &result.Item,
			&result.Status, &result.CreatedAt, &result.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return nil, nil, err
		}
	}

	total, err := repo.countOrder(sqlWhere)
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

func (repo *orderRepository) countOrder(sqlWhere string) (int, error) {
	var count int
	err := repo.db.QueryRow("SELECT COUNT(*) FROM order " + sqlWhere).Scan(&count)
	if err != nil {
		log.Println(err)
		return count, err
	}

	return count, nil
}
