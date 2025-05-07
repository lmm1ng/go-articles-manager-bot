package article

import "go-articles-manager-bot/internal/entities"

type articleRepository interface {
	Create(*entities.Article) error
	GetRandomByTgId(tgId int64) (*entities.Article, error)
	Read(articleId uint32) error
	Delete(articleId uint32) error
	GetArticlesByTgId(tgId int64, read bool, offset uint16, limit uint16) ([]*entities.Article, error)
	GetById(tgId int64, articleId uint32) (*entities.Article, error)
}

type userRepository interface {
	GetByTgUsername(string) (*entities.User, error)
}
