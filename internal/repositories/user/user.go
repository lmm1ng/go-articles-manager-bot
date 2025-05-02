package user

import (
	"database/sql"
	"fmt"
	"time"
)

type UserRepo struct {
	db *sql.DB
}

type User struct {
	ID         uint32
	TgUsername string
	public     bool
	createdAt  time.Time
	updatedAt  time.Time
}

func New(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Prepare() error {
	q := `CREATE TABLE IF NOT EXISTS user (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			tgUsername TEXT NOT NULL UNIQUE,
			public BOOLEAN NOT NULL DEFAULT FALSE,
			createdAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updatedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`
	if _, err := r.db.Exec(q); err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) Create(user *User) error {
	q := `INSERT INTO user (tgUsername, public) VALUES (?, ?)`
	if _, err := r.db.Exec(q, user.TgUsername, user.public); err != nil {
		return fmt.Errorf("Error while creating user, %w", err)
	}

	return nil
}

func (r *UserRepo) GetByUsername(username string) (*User, error) {
	q := `SELECT * FROM user WHERE tgUsername = ?`
	var user User

	if err := r.db.QueryRow(q, username).Scan(&user.ID, &user.TgUsername, &user.public, &user.createdAt, &user.updatedAt); err != nil {
		return nil, fmt.Errorf("Error getting user by tgUsername: %w", err)
	}

	return &user, nil
}

func (r *UserRepo) UpdatePublicByUsername(username string, public bool) error {
	now := time.Now()
	q := `UPDATE user SET public = $2, updatedAt = $3 WHERE tgUsername = $1;`
	if _, err := r.db.Exec(q, username, public, now); err != nil {
		return fmt.Errorf("Error updating user by tgUsername: %w", err)
	}

	return nil
}
