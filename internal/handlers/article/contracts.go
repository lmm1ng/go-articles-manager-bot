package article

import "go-articles-manager-bot/internal/entities"

type articleRepository interface {
	Create(*entities.Article) error
	GetRandomByTgId(tgId int64) (*entities.Article, error)
	Read(articleId uint32) error
	Delete(articleId uint32) error
}

type userRepository interface {
	GetByTgUsername(string) (*entities.User, error)
}
