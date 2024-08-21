package port

type UpdateUserUsecase interface {
	UpdateUserById(command UpdateUserCommand)
}
