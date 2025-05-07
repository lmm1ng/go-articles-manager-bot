package entities

import (
	"fmt"
	"time"
)

type Article struct {
	Id     uint32
	UserId uint32
	Title  *string
	Url    string
	ReadAt *time.Time
}

func (a Article) GetTitleLink() string {
	var text string
	if a.Title != nil {
		text = fmt.Sprintf("[%s](%s)", *a.Title, a.Url)
	} else {
		text = a.Url
	}
	return text
}
