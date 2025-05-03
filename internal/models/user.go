package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id         uint32
	TgUsername string
	Desc       sql.NullString
	TgId       int64
	Public     bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
