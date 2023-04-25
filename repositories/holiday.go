package repositories

import (
	"context"
	"time"
	"vn-lap-interviews/pkg"
)

type HolidayRepository interface {
	IsHoliday(ctx context.Context, date time.Time) bool
	GetEndDate(ctx context.Context, date time.Time) (time.Time, error)
}

type holidayRepository struct {
	db pkg.QueryExecer
}

func NewHolidayRepository(db pkg.QueryExecer) HolidayRepository {
	return &holidayRepository{
		db: db,
	}
}

func (h *holidayRepository) IsHoliday(ctx context.Context, date time.Time) bool {
	query := `
        SELECT is_holiday
        FROM holidays
        WHERE date = $1
    `
	var isHoliday bool
	err := h.db.QueryRow(ctx, query, date.Format(pkg.YYYYMMDD)).Scan(&isHoliday)
	if err != nil {
		return false
	}
	return isHoliday
}

func (h *holidayRepository) GetEndDate(ctx context.Context, date time.Time) (time.Time, error) {

	query := `
		SELECT date
		WHERE date > $1
		AND is_holiday != true
		ORDER by date ASC
		LIMIT 1
	
	`
	var endDate time.Time
	err := h.db.QueryRow(ctx, query, date.Format(pkg.YYYYMMDD)).Scan(&endDate)
	if err != nil {
		return time.Time{}, err
	}
	return endDate, nil
}
