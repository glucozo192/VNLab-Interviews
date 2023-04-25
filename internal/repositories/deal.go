package repositories

import (
	"context"
	"time"
	"vn-lap-interview/internal/entities"
	"vn-lap-interview/pkg"
)

type DealRepository interface {
	GetDeal(ctx context.Context, fromDate, endDate time.Time) ([]*entities.Deal, error)
}

type dealRepository struct {
	db pkg.QueryExecer
}

func NewDealRepository(db pkg.QueryExecer) DealRepository {
	return &dealRepository{
		db: db,
	}
}

func (d *dealRepository) GetDeal(ctx context.Context, fromDate, toDate time.Time) ([]*entities.Deal, error) {
	query := `
		SELECT id, name, due_date, created_at, updated_at
		FROM deals
		WHERE
			due_date > $1 AND due_date < $2
	`
	deals := make([]*entities.Deal, 0)
	rows, err := d.db.Query(ctx, query, fromDate, toDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		deal := new(entities.Deal)
		err := rows.Scan(&deal.ID, &deal.Name, &deal.DueDate, &deal.CreatedAt, &deal.UpdatedAt)
		if err != nil {
			return []*entities.Deal{}, err
		}
		deals = append(deals, deal)
	}
	return deals, nil
}
