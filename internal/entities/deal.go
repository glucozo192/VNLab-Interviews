package entities

import "github.com/jackc/pgtype"

type Deal struct {
	ID        pgtype.Text        `db:"id"`
	Name      pgtype.Text        `db:"name"`
	DueDate   pgtype.Timestamptz `db:"due_date"`
	CreatedAt pgtype.Timestamptz `db:"created_at"`
	UpdatedAt pgtype.Timestamptz `db:"updated_at"`
}
