package scenebuilder

import th "github.com/mymmrac/telego/telegohandler"

type SceneStep struct {
	Cb   th.Handler
	Step uint8
}

func NewSceneStep(
	cb th.Handler,
	step uint8,
) SceneStep {
	return SceneStep{
		Cb:   cb,
		Step: step,
	}
}
