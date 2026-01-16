package user

import "context"

type UserRepository interface {
	Create(ctx context.Context, email, hashPassword string, roleID int) (UserID, error)
	Find(ctx context.Context, id UserID) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
}
