package entities

type User struct {
	Id         uint32
	TgUsername string
	Desc       string
	TgId       int64
	Public     bool
}
