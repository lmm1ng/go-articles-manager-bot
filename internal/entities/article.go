package entities

import (
	"time"
)

type Article struct {
	Id     uint32
	UserId uint32
	Title  *string
	Url    string
	ReadAt *time.Time
}
