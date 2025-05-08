package article

import "go-articles-manager-bot/internal/entities"

type articleRepository interface {
	Create(*entities.Article) error
	GetRandomByTgId(tgId int64) (*entities.Article, error)
	SetRead(articleId uint32, read bool) error
	Delete(articleId uint32) error
	GetArticlesByTgId(tgId int64, read bool, offset uint16, limit uint16) ([]*entities.Article, error)
	GetById(articleId uint32) (*entities.Article, error)
	GetVibe(tgId int64) (*entities.Article, error)
}

type userRepository interface {
	GetByTgId(id int64) (*entities.User, error)
	GetById(id uint32) (*entities.User, error)
}
