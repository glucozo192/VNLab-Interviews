package mocks

import (
	"context"
	"time"

	"vn-lap-interview/internal/entities"

	"github.com/stretchr/testify/mock"
)

type mockDealRepository struct {
	mock.Mock
}

func NewMockDealRepository() *mockDealRepository {
	return &mockDealRepository{}
}

func (m *mockDealRepository) GetDeal(ctx context.Context, fromDate, toDate time.Time) ([]*entities.Deal, error) {
	args := m.Called(fromDate, toDate)
	return args.Get(0).([]*entities.Deal), args.Error(1)
}
