package server

type UserUseCase interface {
	Register()
	Authenticate()
	GetProfile()
}
