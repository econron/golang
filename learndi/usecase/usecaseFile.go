package usecase

import (
	"learndi/repository"
)

type IUsecase interface {
	GetUser() string
}

type Usecase struct {

}

func New() IUsecase {
	return &Usecase{}
}

func (u *Usecase) GetUser() string {
	r := repository.New()
	return r.GetUser()
}