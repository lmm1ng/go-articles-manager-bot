package article

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

type Article struct {
	Id        uint32
	UserId    uint32
	Title     sql.NullString
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
	ReadAt    time.Time
}

func (a *Article) ToEntity() *entities.Article {
	return &entities.Article{
		Id:     a.Id,
		UserId: a.UserId,
		Title:  string(a.Title.String),
		Url:    a.Url,
		ReadAt: a.ReadAt,
	}
}

func (r *repository) Create(article entities.Article) error {
	q := `INSERT INTO article (userId, title, url) VALUES (?, ?, ?)`
	if _, err := r.db.Exec(q, article.UserId, article.Title, article.Url); err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return ErrAlreadyExists
		}
		return fmt.Errorf("Error while creating article, %w", err)
	}

	return nil
}
