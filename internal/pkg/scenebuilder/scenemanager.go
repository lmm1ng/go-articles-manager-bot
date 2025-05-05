package scenebuilder

import "sync"

const ScenesManagerKey = "scenesManager"

const (
	NoScene uint8 = iota
	// create article
	StepAddArticleUrl
	// set user desc
	StepAddUserDesc
)

type SceneManager struct {
	Users *map[int64]uint8
	Mutex *sync.RWMutex
}

func NewSceneManager(users *map[int64]uint8, mutex *sync.RWMutex) *SceneManager {
	return &SceneManager{
		Users: users,
		Mutex: mutex,
	}
}
