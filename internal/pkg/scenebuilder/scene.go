package scenebuilder

import (
	"errors"
	"go-articles-manager-bot/internal/handlers"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type Scene struct {
	steps     []SceneStep
	predicate th.Predicate
}

func NewScene(
	steps []SceneStep,
	predicate th.Predicate,
) Scene {
	return Scene{
		steps:     steps,
		predicate: predicate,
	}
}

func (s *Scene) Register(
	h *th.BotHandler,
) []handlers.Handler {
	outSteps := make([]handlers.Handler, len(s.steps))
	for pos, step := range s.steps {
		f := func(ctx *th.Context, update telego.Update) error {
			scenesManager := ctx.Value(ScenesManagerKey).(*SceneManager)
			if scenesManager == nil {
				return errors.New("No users state")
			}
			scenesManager.Mutex.RLock()
			curScene, exists := scenesManager.Users[update.Message.From.ID]
			scenesManager.Mutex.RUnlock()

			if !exists {
				return ctx.Next(update)
			}

			if curScene != uint8(step.Step) {
				return ctx.Next(update)
			}

			err := step.Cb(ctx, update)

			if err != nil {
				return nil
			}

			scenesManager.Mutex.Lock()
			if pos == len(s.steps)-1 {
				scenesManager.Users[update.Message.From.ID] = NoScene
			} else {
				scenesManager.Users[update.Message.From.ID] = s.steps[pos+1].Step
			}

			scenesManager.Mutex.Unlock()

			return nil
		}

		var pred th.Predicate

		if pos == 0 {
			pred = s.predicate
		} else {
			pred = th.Any()
		}
		outSteps[pos] = handlers.NewHandler(f, pred)
	}

	for _, step := range outSteps {
		h.Handle(step.Cb, step.Predicate)
	}

	return nil
}
