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

func (a *Article) toEntity() *entities.Article {
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
	q := `
	SELECT * FROM article
	WHERE userId in (SELECT id FROM user WHERE tgId = ?) AND
	readAt IS NULL
	ORDER BY RANDOM()
	LIMIT 1`
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

	return article.toEntity(), nil
}

func (r *repository) GetVibe(tgId int64) (*entities.Article, error) {
	q := `
	SELECT * FROM article
	WHERE userId in (SELECT id FROM user WHERE tgId != ? AND public = TRUE ORDER BY RANDOM())
	ORDER BY RANDOM()
	LIMIT 1`

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
		return nil, fmt.Errorf("Error while getting vibe article, %w", err)
	}

	return article.toEntity(), nil
}

func (r *repository) GetById(articleId uint32) (*entities.Article, error) {
	q := `SELECT * FROM article WHERE article.id = ?;`
	var article Article
	row := r.db.QueryRow(q, articleId)

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
		return nil, fmt.Errorf("Error while getting article, %w", err)
	}

	return article.toEntity(), nil
}

func (r *repository) SetRead(articleId uint32, read bool) error {
	t := sql.NullTime{Valid: read, Time: time.Now()}

	q := `UPDATE article SET readAt = ?, updatedAt = DATETIME() WHERE id = ?;`

	row, err := r.db.Exec(q, t, articleId)

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
func (r *repository) GetArticlesByTgId(tgId int64, read bool, offset uint16, limit uint16) ([]*entities.Article, error) {
	readAt := sql.NullTime{Valid: read, Time: time.Now()}

	q := `
	SELECT a.* FROM article a
	WHERE a.userId in (SELECT u.id FROM user u WHERE u.tgId = ?) AND
	(a.readAt < ? OR a.readAt IS NULL)
	ORDER BY a.id ASC
	LIMIT ?
	OFFSET ?`

	rows, err := r.db.Query(q, tgId, readAt, limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var articles []*entities.Article

	for rows.Next() {
		var a Article
		if err := rows.Scan(
			&a.Id,
			&a.UserId,
			&a.Title,
			&a.Url,
			&a.CreatedAt,
			&a.UpdatedAt,
			&a.ReadAt,
		); err != nil {
			return articles, err
		}

		articles = append(articles, a.toEntity())
	}

	return articles, nil
}

func (r *repository) GetArticlesCountByPeriod(tgId int64, read bool, start time.Time, end time.Time) (uint16, error) {
	readAt := sql.NullTime{Valid: read, Time: time.Now()}

	var count uint16

	q := `
	SELECT COUNT(*) FROM article
	WHERE userId in (SELECT id FROM user WHERE tgId = ?) AND
	(readAt < ? OR readAt IS NULL) AND
	updatedAt < ? AND updatedAt > ?;`

	row := r.db.QueryRow(q, tgId, readAt, end, start)

	if err := row.Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return count, ErrNotFound
		}
		return count, fmt.Errorf("Error while getting articles count, %w", err)
	}

	return count, nil
}
