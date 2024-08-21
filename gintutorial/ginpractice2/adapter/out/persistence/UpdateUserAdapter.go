package persistence

import (
	"context"
	"ginpractice2/adapter/out/persistence/mysql"
	model "ginpractice2/application/domain/model"
)

type UpdateUserAdapter struct {
	Queries *mysql.Queries
}

func (a *UpdateUserAdapter) UpdateUser(userId int64, name string) bool {
	arg := mysql.UpdateUserParams{
		Name: name,
		ID: userId,
	}

	err := a.Queries.UpdateUser(context.TODO(), arg)

	if err != nil { return false }; return true

}

func (a *UpdateUserAdapter) GetUser(userId int64) (*model.User, error) {
	user := model.User{}

	ret, err := a.Queries.FindUserById(context.TODO(), userId)
	if err != nil {
		
		return &user, err
	}

	user.UserID = ret.ID
	user.Name = ret.Name

	return &user, nil
}