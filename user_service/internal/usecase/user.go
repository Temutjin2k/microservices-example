package usecase

type UserUseCase struct {
	userRepo UserRepo
}

func NewUser(userRepo UserRepo) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}
