package repository

import (
	"context"
	"errors"

	"github.com/cortezzIP/realtime-leaderboard-api/internal/database/postgres"
	"github.com/cortezzIP/realtime-leaderboard-api/internal/model"
	"github.com/jackc/pgx/v5"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	GetUserById(ctx context.Context, id int) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateRating(ctx context.Context, id int, rating int) error
}

type UserRepo struct {
	db *pgx.Conn
}

func NewUserRepository() *UserRepo {
	return &UserRepo{
		db: postgres.GetDatabase(),
	}
}

func (r *UserRepo) GetUserById(ctx context.Context, id int) (*model.User, error) {
	query := "SELECT * FROM users WHERE id = $1"

	var user model.User

	err := r.db.QueryRow(ctx, query, id).Scan(&user.Id, &user.Login, &user.Password, &user.Rating)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	query := "SELECT * FROM users WHERE username = $1"

	var user model.User

	err := r.db.QueryRow(ctx, query, username).Scan(&user.Id, &user.Login, &user.Password, &user.Rating)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) CreateUser(ctx context.Context, user *model.User) error {
	query := "INSERT INTO users (login, password, rating) VALUES $1, $2, $3"

	_, err := r.db.Exec(ctx, query, user.Login, user.Password, user.Rating)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) UpdateRating(ctx context.Context, id int, rating int) error {
	query := "UPDATE users SET rating = $1 WHERE id = $2"

	commandTag, err := r.db.Exec(ctx, query, id, rating)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return ErrUserNotFound
	}

	return nil
}