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
	users map[int64]uint8
	mutex sync.Mutex
}

func NewSceneManager() *SceneManager {
	m := make(map[int64]uint8)
	return &SceneManager{
		users: m,
	}
}
