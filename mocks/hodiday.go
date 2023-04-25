package mocks

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type mockHolidayRepository struct {
	mock.Mock
}

func NewMockHolidayRepository() *mockHolidayRepository {
	return &mockHolidayRepository{}
}
func (m *mockHolidayRepository) IsHoliday(ctx context.Context, date time.Time) (bool, error) {
	args := m.Called(date)
	return args.Bool(0), args.Error(1)
}

func (m *mockHolidayRepository) GetEndDate(ctx context.Context, date time.Time) (time.Time, error) {
	args := m.Called(date)
	return args.Get(0).(time.Time), args.Error(1)
}
