package handlers

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type Cb = func(ctx *th.Context, update telego.Update) error

type Handler struct {
	Cb        Cb
	Predicate th.Predicate
}

func New(cb Cb, predicate th.Predicate) Handler {
	return Handler{
		Cb:        cb,
		Predicate: predicate,
	}
}
