package article

import (
	"database/sql"
	"fmt"
	"time"
)

type ArticleRepo struct {
	db *sql.DB
}

type Article struct {
	id        uint32
	userId    uint32
	title     string
	url       string
	createdAt time.Time
	updatedAt time.Time
	readAt    time.Time
}

func New(db *sql.DB) *ArticleRepo {
	return &ArticleRepo{
		db: db,
	}
}

func (r *ArticleRepo) Prepare() error {
	q := `CREATE TABLE IF NOT EXISTS article (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			userId INTEGER NOT NULL,
			title TEXT NOT NULL,
			url TEXT NOT NULL,
			createdAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updatedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			readAt DATETIME DEFAULT NULL,
			FOREIGN KEY (userId) REFERENCES user(id) ON DELETE CASCADE
		);`
	if _, err := r.db.Exec(q); err != nil {
		return err
	}

	return nil
}

func (r *ArticleRepo) Create(article *Article) error {
	q := `INSERT INTO article (userId, title, url) VALUES (?, ?, ?)`
	if _, err := r.db.Exec(q, article.userId, article.title, article.url); err != nil {
		return fmt.Errorf("Error while creating article, %w", err)
	}

	return nil
}
