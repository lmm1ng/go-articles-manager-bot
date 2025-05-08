package user

import "strings"

type UserHandler struct {
	userRepo userRepository
}

func New(userRepo userRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
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
