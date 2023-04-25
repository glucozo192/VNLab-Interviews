package domain

import (
	"context"
	"fmt"
	"time"
	"vn-lap-interviews/entities"
	"vn-lap-interviews/repositories"
)

type SendMailDomain interface {
	SendMail(context.Context) string
}

type sendMailDomain struct {
	holidayRepo repositories.HolidayRepository
	dealRepo    repositories.DealRepository
}

func NewSendMailDomain() SendMailDomain {
	return &sendMailDomain{}
}

// todo: iSHoliday func should recieve multiple values

func (s *sendMailDomain) SendMail(ctx context.Context) string {
	currentDate := time.Now()
	if s.isHoliday(ctx, currentDate) {
		return "skip holiday"
	}

	nextThreeDate := time.Now().Add(3 * 24 * time.Hour)
	if s.isHoliday(ctx, nextThreeDate) {
		return "skip holiday"
	}

	endDate, err := s.endDateSendMail(ctx, nextThreeDate)
	if err != nil {
		return "can't get end date"
	}

	deals, err := s.getDeal(ctx, nextThreeDate, endDate)
	if err != nil {
		return "can't get deal"
	}

	for deal := range deals {
		fmt.Println("sendmail: ", deal)
	}

	return "success"
}

func (s *sendMailDomain) isHoliday(ctx context.Context, date ...time.Time) bool {
	return s.holidayRepo.IsHoliday(ctx, date)
}

func (s *sendMailDomain) endDateSendMail(ctx context.Context, date time.Time) (time.Time, error) {
	endDate, err := s.holidayRepo.GetEndDate(ctx, date)
	if err != nil {
		return time.Time{}, err
	}
	return endDate, nil
}

func (s *sendMailDomain) getDeal(ctx context.Context, fromDate, toDate time.Time) ([]*entities.Deal, error) {
	deals, err := s.dealRepo.GetDeal(ctx, fromDate, toDate)
	if err != nil {
		return nil, err
	}
	return deals, nil
}
