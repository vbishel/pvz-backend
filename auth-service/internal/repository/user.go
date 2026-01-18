package repository

import (
	"auth-service/internal/domain/role"
	"auth-service/internal/domain/user"
	"auth-service/pkg/apperrors"
	"auth-service/pkg/database/postgres"
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
)

const usersTable = "users"

type userRepository struct {
	*postgres.Postgres
}

func NewUserRepository(pg *postgres.Postgres) user.UserRepository {
	return &userRepository{pg}
}

func (r *userRepository) Create(ctx context.Context, email, hashPassword string) (user.UserID, error) {
	sql, args, err := r.Builder.
		Insert(usersTable).
		Columns("email, password, role_id").
		Values(email, hashPassword, role.RoleClientId).
		Suffix("RETURNING ID").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("r.Builder.Insert: %w", err)
	}

	var uid user.UserID

	if err = r.Pool.QueryRow(ctx, sql, args...).Scan(&uid); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return 0, fmt.Errorf("r.Pool.Exec: %w", apperrors.ErrUserAlreadyExists)
		}

		return 0, fmt.Errorf("r.Pool.Exec: %w", err)
	}

	return uid, nil
}

func (r *userRepository) Find(ctx context.Context, id user.UserID) (user.User, error) {
	sql, args, err := r.Builder.
		Select("email, password").
		From(usersTable).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return user.User{}, fmt.Errorf("r.Builder.Select: %w", err)
	}

	u := user.User{ID: id}

	if err = r.Pool.QueryRow(ctx, sql, args...).Scan(
		&u.Email,
		&u.Password,
	); err != nil {
		if err == pgx.ErrNoRows {
			return user.User{}, fmt.Errorf("r.Pool.QueryRow.Scan: %w", apperrors.ErrUserNotFound)
		}

		return user.User{}, fmt.Errorf("r.Pool.QueryRow.Scan: %w", err)
	}

	return u, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (user.User, error) {
	sql, args, err := r.Builder.
		Select("id, password").
		From(usersTable).
		Where(squirrel.Eq{"email": email}).
		ToSql()
	if err != nil {
		return user.User{}, fmt.Errorf("r.Builder.Select: %w", err)
	}

	u := user.User{Email: email}

	if err = r.Pool.QueryRow(ctx, sql, args...).Scan(
		&u.ID,
		&u.Password,
	); err != nil {
		if err == pgx.ErrNoRows {
			return user.User{}, fmt.Errorf("r.Pool.QueryRow.Scan: %w", apperrors.ErrUserNotFound)
		}

		return user.User{}, fmt.Errorf("r.Pool.QueryRow.Scan: %w", err)
	}

	return u, nil
}

