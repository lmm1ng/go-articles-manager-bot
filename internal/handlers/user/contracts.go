package user

import "go-articles-manager-bot/internal/entities"

type userRepository interface {
	Create(user *entities.User) error
	UpdateDescByTgId(id int64, desc string) error
	UpdatePublicByTgId(id int64, public bool) error
	GetByTgId(id int64) (*entities.User, error)
}
