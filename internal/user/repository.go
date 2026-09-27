package user

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) AddUser(ctx context.Context, chatId int, info any) error {
	query := `
		INSERT INTO bot_users (chat_id, info)
		VALUES ($1, $2)
	`

	_, err := r.db.Exec(ctx, query, info)

	return err
}

func (r *Repository) GetUser(ctx context.Context, id int64) (*User, error) {
	var u User
	err := s.db.GetContext(ctx, &u, `SELECT id, info, created_at, mailing, sent_lead FROM bot_users WHERE id = $1`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

type UserStorage interface {
	AddUser(ctx context.Context, u *User) (int64, error)
	GetUser(ctx context.Context, id int64) (*User, error)
}
