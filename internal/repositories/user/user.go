package user

import (
	"database/sql"
	"fmt"
	"go-articles-manager-bot/internal/models"
	"time"
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(user *models.User) error {
	q := `INSERT INTO user (tgUsername, tgId) VALUES (?, ?)`
	fmt.Println("kekeke")
	if _, err := r.db.Exec(q, user.TgUsername, user.TgId); err != nil {
		fmt.Println("Error while creating user:", err)
		return fmt.Errorf("Error while creating user, %w", err)
	}

	return nil
}

func (r *repository) GetByTgUsername(username string) (*models.User, error) {
	q := `SELECT * FROM user WHERE tgUsername = ?`
	var user models.User

	if err := r.db.QueryRow(q, username).Scan(&user.Id, &user.TgId, &user.TgUsername, &user.Desc, &user.Public, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, fmt.Errorf("Error getting user by tgUsername: %w", err)
	}

	return &user, nil
}

func (r *repository) UpdatePublicByUsername(username string, public bool) error {
	now := time.Now()
	q := `UPDATE user SET public = $2, updatedAt = $3 WHERE tgUsername = $1;`
	if _, err := r.db.Exec(q, username, public, now); err != nil {
		return fmt.Errorf("Error updating user by tgUsername: %w", err)
	}

	return nil
}
