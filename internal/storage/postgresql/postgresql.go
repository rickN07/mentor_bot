package storage

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rickN07/mentor_bot/internal/config"
	"github.com/rickN07/mentor_bot/internal/e"
	"github.com/rickN07/mentor_bot/internal/utils"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewPostgresClient(logger *slog.Logger, ctx context.Context, config config.DB, maxAttempts int) (pool *pgxpool.Pool, err error) {
	defer func() { err = e.WrapIfErr("can't get updates", err) }()

	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		password := url.QueryEscape(config.Password)
		dsn := fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s",
			config.User,
			password,
			config.Host,
			config.Port,
			config.Database,
		)

		cfg, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			logger.Error("can't parse config by pgxpool", "err", err)
			return err
		}

		pool, err = pgxpool.NewWithConfig(ctx, cfg)
		if err != nil {
			logger.Error("can't make a new connection with config by pgxpool", "err", err)
			return err
		}

		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {
		logger.Error("can't make a new storage client", "err", err)
		return nil, err
	}

	return pool, nil
}
