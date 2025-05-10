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
			var from int64

			if update.Message == nil {
				from = update.CallbackQuery.From.ID
			} else {
				from = update.Message.From.ID
			}

			scenesManager := ctx.Value(ScenesManagerKey).(*SceneManager)
			if scenesManager == nil {
				return errors.New("No users state")
			}
			scenesManager.mutex.Lock()
			curScene := scenesManager.users[from]
			scenesManager.mutex.Unlock()

			if curScene != uint8(step.Step) {
				return ctx.Next(update)
			}

			err := step.Cb(ctx, update)

			if err != nil {
				return nil
			}

			scenesManager.mutex.Lock()
			if pos == len(s.steps)-1 {
				scenesManager.users[from] = NoScene
			} else {
				scenesManager.users[from] = s.steps[pos+1].Step
			}

			scenesManager.mutex.Unlock()

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
