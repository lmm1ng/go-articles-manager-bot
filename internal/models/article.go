package models

import "time"

type Article struct {
	Id        uint32
	UserId    uint32
	Title     string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
	ReadAt    time.Time
}
