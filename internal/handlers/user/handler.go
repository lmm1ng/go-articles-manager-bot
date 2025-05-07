package user

type UserHandler struct {
	userRepo userRepository
}

func New(userRepo userRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}
