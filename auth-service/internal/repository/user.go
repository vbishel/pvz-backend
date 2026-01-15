package repository

import (
	"auth-service/internal/domain/entity"
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

func NewUserRepository(pg *postgres.Postgres) *userRepository {
	return &userRepository{pg}
}

func (r *userRepository) Create(ctx context.Context, u entity.User) (entity.UserID, error) {
	sql, args, err := r.Builder.
		Insert(usersTable).
		Columns("email, password, role_id").
		Values(u.Email, u.Password, u.RoleID).
		Suffix("RETURNING ID").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("r.Builder.Insert: %w", err)
	}

	var uid entity.UserID

	if err = r.Pool.QueryRow(ctx, sql, args...).Scan(&uid); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return 0, fmt.Errorf("r.Pool.Exec: %w", apperrors.ErrUserAlreadyExists)
		}

		return 0, fmt.Errorf("r.Pool.Exec: %w", err)
	}

	return uid, nil
}

func (r *userRepository) Find(ctx context.Context, id entity.UserID) (entity.User, error) {
	sql, args, err := r.Builder.
		Select("email, password").
		From(usersTable).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("r.Builder.Select: %w", err)
	}

	user := entity.User{ID: id}

	if err = r.Pool.QueryRow(ctx, sql, args...).Scan(
		&user.Email,
		&user.Password,
	); err != nil {
		if err == pgx.ErrNoRows {
			return entity.User{}, fmt.Errorf("r.Pool.QueryRow.Scan: %w", apperrors.ErrUserNotFound)
		}

		return entity.User{}, fmt.Errorf("r.Pool.QueryRow.Scan: %w", err)
	}

	return user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	sql, args, err := r.Builder.
		Select("id, password").
		From(usersTable).
		Where(squirrel.Eq{"email": email}).
		ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("r.Builder.Select: %w", err)
	}

	user := entity.User{Email: email}

	if err = r.Pool.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Password,
	); err != nil {
		if err == pgx.ErrNoRows {
			return entity.User{}, fmt.Errorf("r.Pool.QueryRow.Scan: %w", apperrors.ErrUserNotFound)
		}

		return entity.User{}, fmt.Errorf("r.Pool.QueryRow.Scan: %w", err)
	}

	return user, nil
}


