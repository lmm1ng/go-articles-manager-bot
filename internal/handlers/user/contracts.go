package user

import (
	"go-articles-manager-bot/internal/entities"
	"time"
)

type userRepository interface {
	Create(user *entities.User) error
	UpdateDescByTgId(id int64, desc string) error
	UpdatePublicByTgId(id int64, public bool) error
	GetByTgId(id int64) (*entities.User, error)
}

type articleRepository interface {
	GetArticlesCountByPeriod(tgId int64, read bool, start time.Time, end time.Time) (uint16, error)
}
