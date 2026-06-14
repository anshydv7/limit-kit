package repository

import (
	"auth-service/base/database"
	"auth-service/models/domain"
	"context"
	"time"
)

func GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	db := database.GetDB()
	query := `SELECT id, email, name, profile_picture, deleted_at FROM users WHERE email = $1 AND deleted_at IS NULL`

	row := db.QueryRowContext(ctx, query, email)
	var user domain.User
	if err := row.Scan(&user.ID, &user.Email, &user.Name, &user.ProfilePicture, &user.DeletedAt); err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	db := database.GetDB()
	query := `INSERT INTO users (email, name, profile_picture, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	err := db.QueryRowContext(ctx, query, user.Email, user.Name, user.ProfilePicture, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
