package port

import (
	model "ginpractice2/application/domain/model"
)

type GetUserPort interface {
	GetUser(userId string) (*model.User, error)
}