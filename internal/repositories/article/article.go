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
	ReadAt    sql.NullTime
}

func (a *Article) ToEntity() *entities.Article {
	var readAt *time.Time
	if a.ReadAt.Valid {
		readAt = &a.ReadAt.Time
	}
	return &entities.Article{
		Id:     a.Id,
		UserId: a.UserId,
		Title:  a.Title.String,
		Url:    a.Url,
		ReadAt: readAt,
	}
}

func (r *repository) Create(article *entities.Article) error {
	q := `INSERT INTO article (userId, title, url) VALUES (?, ?, ?)`
	if _, err := r.db.Exec(q, article.UserId, article.Title, article.Url); err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return ErrAlreadyExists
		}
		return fmt.Errorf("Error while creating article, %w", err)
	}

	return nil
}

func (r *repository) GetRandomByTgId(tgId int64) (*entities.Article, error) {
	q := `SELECT a.* FROM article a WHERE a.userId in (SELECT u.id FROM user u WHERE u.tgId = ?) ORDER BY RANDOM() LIMIT 1;`
	var article Article
	row := r.db.QueryRow(q, tgId)

	if err := row.Scan(
		&article.Id,
		&article.UserId,
		&article.Title,
		&article.Url,
		&article.CreatedAt,
		&article.UpdatedAt,
		&article.ReadAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("Error while getting random article, %w", err)
	}

	return article.ToEntity(), nil
}
