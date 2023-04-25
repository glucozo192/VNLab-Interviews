package repositories

import (
	"context"
	"time"
	"vn-lap-interviews/entities"
	"vn-lap-interviews/pkg"
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
	return []*entities.Deal{}, nil
}
