package port

import (
	model "ginpractice2/application/domain/model"
)

type UpdateUserPort interface {
	UpdateUser(u *model.User) bool
}