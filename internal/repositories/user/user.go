package user

import (
	"database/sql"
	"fmt"
	"go-articles-manager-bot/internal/entities"
	"strings"
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

type User struct {
	Id         uint32
	TgUsername string
	Desc       string
	TgId       int64
	Public     bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (u *User) ToEntity() *entities.User {
	return &entities.User{
		Id:         u.Id,
		TgUsername: u.TgUsername,
		Desc:       u.Desc,
		TgId:       u.TgId,
		Public:     u.Public,
	}
}

func (r *repository) Create(user *entities.User) error {
	q := `INSERT INTO user (tgUsername, tgId) VALUES (?, ?)`
	if _, err := r.db.Exec(q, user.TgUsername, user.TgId); err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return ErrAlreadyExists
		}
		return fmt.Errorf("Error while creating user, %w", err)
	}

	return nil
}

func (r *repository) GetByTgUsername(username string) (*entities.User, error) {
	q := `SELECT * FROM user WHERE tgUsername = ?`
	var user User

	if err := r.db.QueryRow(q, username).Scan(&user.Id, &user.TgId, &user.TgUsername, &user.Desc, &user.Public, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, ErrNotFound
	}

	return user.ToEntity(), nil
}

func (r *repository) UpdatePublicByUsername(username string, public bool) error {
	now := time.Now()
	q := `UPDATE user SET public = $2, updatedAt = $3 WHERE tgUsername = $1;`
	if _, err := r.db.Exec(q, username, public, now); err != nil {
		return fmt.Errorf("Error updating user by tgUsername: %w", err)
	}

	return nil
}
