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
	TgUsername sql.NullString
	Desc       sql.NullString
	TgId       int64
	Public     bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (u *User) ToEntity() *entities.User {
	var desc *string
	var username *string

	if u.Desc.Valid {
		desc = &u.Desc.String
	}

	if u.TgUsername.Valid {
		username = &u.TgUsername.String
	}
	return &entities.User{
		Id:         u.Id,
		TgUsername: username,
		Desc:       desc,
		TgId:       u.TgId,
		Public:     u.Public,
	}
}

func (r *repository) Create(user *entities.User) error {
	username := sql.NullString{Valid: false, String: ""}

	if user.TgUsername != nil {
		username = sql.NullString{Valid: true, String: *user.TgUsername}
	}

	q := `INSERT INTO user (tgUsername, tgId) VALUES (?, ?)`

	if _, err := r.db.Exec(q, username, user.TgId); err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return ErrAlreadyExists
		}
		return fmt.Errorf("Error while creating user, %w", err)
	}

	return nil
}

func (r *repository) GetByTgId(id int64) (*entities.User, error) {
	q := `SELECT * FROM user WHERE tgId = ?`
	var user User

	if err := r.db.QueryRow(q, id).Scan(
		&user.Id,
		&user.TgId,
		&user.TgUsername,
		&user.Desc,
		&user.Public,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, ErrNotFound
	}

	return user.ToEntity(), nil
}

func (r *repository) GetById(id uint32) (*entities.User, error) {
	q := `SELECT * FROM user WHERE id = ?`
	var user User

	if err := r.db.QueryRow(q, id).Scan(
		&user.Id,
		&user.TgId,
		&user.TgUsername,
		&user.Desc,
		&user.Public,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, ErrNotFound
	}

	return user.ToEntity(), nil
}

func (r *repository) UpdatePublicByTgId(id int64, public bool) error {
	q := `UPDATE user SET public = ?, updatedAt = DATE() WHERE tgId = ?;`
	row, err := r.db.Exec(q, public, id)

	if err != nil {
		return fmt.Errorf("Error updating user public by telegram id: %w", err)
	}

	if c, _ := row.RowsAffected(); c == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *repository) UpdateDescByTgId(id int64, desc string) error {
	q := `UPDATE user SET desc = ?, updatedAt = DATE() WHERE tgId = ?;`
	row, err := r.db.Exec(q, desc, id)

	if err != nil {
		return fmt.Errorf("Error updating user desc by telegram id: %w", err)
	}

	if c, _ := row.RowsAffected(); c == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *repository) GetDescByTgId(id int64, desc string) error {
	q := `UPDATE user SET desc = ?, updatedAt = DATE() WHERE tgId = ?;`
	row, err := r.db.Exec(q, desc, id)

	if err != nil {
		return fmt.Errorf("Error updating user desc by telegram id: %w", err)
	}

	if c, _ := row.RowsAffected(); c == 0 {
		return ErrNotFound
	}

	return nil
}
