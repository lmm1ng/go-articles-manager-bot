package user

import "strings"

type UserHandler struct {
	userRepo    userRepository
	articleRepo articleRepository
}

func New(userRepo userRepository, articleRepo articleRepository) *UserHandler {
	return &UserHandler{
		userRepo:    userRepo,
		articleRepo: articleRepo,
	}
}

func (uh *UserHandler) getCallbackArgs(cbStr string, buttons ...string) []string {
	argsStr := strings.Clone(cbStr)

	for _, b := range buttons {
		s, found := strings.CutPrefix(argsStr, b+" ")
		if !found {
			continue
		}
		argsStr = s
	}

	return strings.Split(argsStr, " ")
}
