package article

import (
	"database/sql"
	"fmt"
	"go-articles-manager-bot/internal/models"
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Prepare() error {
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

func (r *repository) Create(article *models.Article) error {
	q := `INSERT INTO article (userId, title, url) VALUES (?, ?, ?)`
	if _, err := r.db.Exec(q, article.UserId, article.Title, article.Url); err != nil {
		fmt.Println(err)
		return fmt.Errorf("Error while creating article, %w", err)
	}

	return nil
}
