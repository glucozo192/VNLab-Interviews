package domain

import (
	"context"
	"testing"
	"time"
	"vn-lap-interviews/internal/entities"
	"vn-lap-interviews/mocks"

	"github.com/stretchr/testify/mock"
)

func Test_sendMailDomain_SendMail(t *testing.T) {
	holidayRepo := mocks.NewMockHolidayRepository()
	dealRepo := mocks.NewMockDealRepository()
	domain := NewSendMailDomain(holidayRepo, dealRepo)
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name  string
		args  args
		want  error
		setup func(context.Context)
	}{
		{
			name: "happy case",
			setup: func(ctx context.Context) {
				holidayRepo.On("IsHoliday", mock.Anything, mock.Anything).Times(2).Return(false, nil)
				holidayRepo.On("GetEndDate", mock.Anything, mock.Anything).Once().Return(time.Time{}, nil)
				dealRepo.On("GetDeal", mock.Anything, mock.Anything, mock.Anything).Once().Return([]*entities.Deal{}, nil)
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(context.Background())
			if got := domain.SendMail(tt.args.ctx); got != tt.want {
				t.Errorf("sendMailDomain.SendMail() = %v want %v", got, tt.want)
			}
		})
	}
}
