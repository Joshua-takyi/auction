package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id" json:"id"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	Password  string    `db:"password" json:"password"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type UserRepo interface {
	CreateUser(ctx context.Context, user User) (User, error)

	UpdateUser(ctx context.Context, user_id uuid.UUID, update map[string]any) (User, error)
	DeleteUser(ctx context.Context, user_id uuid.UUID) (string, error)
	ChangePassword(ctx context.Context, user_id uuid.UUID, newPassword string) (string, error)
}

// func () CreateUser(ctx context.Context, user User) (User, error) {
// 	return nil, nil
// }
