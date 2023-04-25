package domain

import (
	"context"
	"fmt"
	"time"
	"vn-lap-interviews/internal/entities"
	"vn-lap-interviews/internal/repositories"
)

type SendMailDomain interface {
	SendMail(context.Context) error
}

type sendMailDomain struct {
	holidayRepo repositories.HolidayRepository
	dealRepo    repositories.DealRepository
}

func NewSendMailDomain(
	holidayRepo repositories.HolidayRepository,
	dealRepo repositories.DealRepository,
) SendMailDomain {
	return &sendMailDomain{holidayRepo, dealRepo}
}

// todo: iSHoliday func should recieve multiple values

func (s *sendMailDomain) SendMail(ctx context.Context) error {
	currentDate := time.Now()
	isHoliday, err := s.isHoliday(ctx, currentDate)
	if err != nil {
		return err
	}
	if isHoliday {
		return nil
	}

	nextThreeDate := time.Now().Add(3 * 24 * time.Hour)
	isHoliday, err = s.isHoliday(ctx, nextThreeDate)
	if err != nil {
		return err
	}
	if isHoliday {
		return nil
	}

	endDate, err := s.endDateSendMail(ctx, nextThreeDate)
	if err != nil {
		return fmt.Errorf("can't get end date: %v", err)
	}

	deals, err := s.getDeal(ctx, nextThreeDate, endDate)
	if err != nil {
		return fmt.Errorf("can't get deal: %v", err)
	}

	for deal := range deals {
		fmt.Println("send mail: ", deal)
	}

	return nil
}

func (s *sendMailDomain) isHoliday(ctx context.Context, date time.Time) (bool, error) {
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
