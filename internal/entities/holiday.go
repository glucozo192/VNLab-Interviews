package entities

import "github.com/jackc/pgtype"

type Holiday struct {
	ID        pgtype.Text        `db:"id"`
	Date      pgtype.Timestamptz `db:"date"`
	IsHoliday pgtype.Bool        `db:"is_holiday"`
	CreatedAt pgtype.Timestamptz `db:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at"`
}
