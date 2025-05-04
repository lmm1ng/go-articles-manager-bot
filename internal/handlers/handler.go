package handlers

import (
	th "github.com/mymmrac/telego/telegohandler"
)

type Handler struct {
	Cb        th.Handler
	Predicate th.Predicate
}

func NewHandler(cb th.Handler, predicate th.Predicate) Handler {
	return Handler{
		Cb:        cb,
		Predicate: predicate,
	}
}
