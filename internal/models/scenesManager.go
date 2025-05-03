package models

import "sync"

type ScenesManager struct {
	Users map[int64]uint8
	Mutex sync.RWMutex
}

const (
	NoScene uint8 = iota
	StateAddUrl
)
