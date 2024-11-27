package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	
	"github.com/cortezzIP/realtime-leaderboard-api/internal/config"
)

var db *pgx.Conn

func Connect(ctx context.Context, cfg *config.PostgresConfig) error {
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	var err error

	db, err = pgx.Connect(ctx, databaseUrl)
	if err != nil {
		slog.Error("аailed to connect to the database: " + err.Error())
		return err
	}

	err = db.Ping(ctx)
	if err != nil {
		slog.Error("аailed to connect to ping database: " + err.Error())
		return err
	}

	return nil
}

func GetDatabase() *pgx.Conn {
	return db
}

func Close() {
	if db != nil {
		db.Close(context.TODO())
	}
}