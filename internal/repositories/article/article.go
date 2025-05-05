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
	var title *string

	if a.ReadAt.Valid {
		readAt = &a.ReadAt.Time
	}

	if a.Title.Valid {
		title = &a.Title.String
	}

	return &entities.Article{
		Id:     a.Id,
		UserId: a.UserId,
		Title:  title,
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
	q := `SELECT a.* FROM article a WHERE a.userId in (SELECT u.id FROM user u WHERE u.tgId = ?) AND a.readAt IS NULL ORDER BY RANDOM() LIMIT 1;`
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

func (r *repository) Read(articleId uint32) error {
	q := `UPDATE article SET readAt = DATE() WHERE id = ?`

	row, err := r.db.Exec(q, articleId)

	if err != nil {
		return fmt.Errorf("Error while reading article, %w", err)
	}

	if c, _ := row.RowsAffected(); c == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *repository) Delete(articleId uint32) error {
	q := `DELETE FROM article WHERE id = ?`
	row, err := r.db.Exec(q, articleId)

	if err != nil {
		return fmt.Errorf("Error while deleting article, %w", err)
	}

	if c, _ := row.RowsAffected(); c == 0 {
		return ErrNotFound
	}

	return nil
}
func (r *repository) ShowArticlesByTgId(tgId int64, readed bool, offset uint16) (*entities.Article, error) {
	limit := 10
	var t time.Time
	if readed {
		t = time.Now()
	}

	q := `
	SELECT a.* FROM article a
	WHERE a.userId in (SELECT u.id FROM user u WHERE u.tgId = $1) AND
	a.readAt < $2
	ORDER BY a.id DESC
	LIMIT 10
	OFFSET $3;`

	var article Article
	rows, err := r.db.Query(q, tgId)

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
