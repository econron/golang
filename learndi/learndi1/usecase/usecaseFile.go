package usecase

import (
	"learndi/repository"
)

type IUsecase interface {
	GetUser() string
}

type Usecase struct {
	repo repository.IRepository
}

func New(r repository.IRepository) IUsecase {
	return &Usecase{
		repo: r,
	}
}

func (u *Usecase) GetUser() string {
	return u.repo.GetUser()
}