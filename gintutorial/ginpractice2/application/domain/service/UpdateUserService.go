package service

import(
	"ginpractice2/application/port/in"
	"ginpractice2/adapter/out/persistence"
	model "ginpractice2/application/domain/model"
)

type UpdateUserService struct {
	A *persistence.UpdateUserAdapter
}

func (s *UpdateUserService) UpdateUserById(c *port.UpdateUserCommand) bool {
	_, err := s.getUserById(c.UserID)
	if err != nil {
		return false
	}
	return s.A.UpdateUser(c.UserID, c.Name)
}

func (s *UpdateUserService) getUserById(userId int64) (*model.User, error) {
	ret, err := s.A.GetUser(userId)
	if err != nil {
		return &model.User{}, err
	}
	return ret, nil
}