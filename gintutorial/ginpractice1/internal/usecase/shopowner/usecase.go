package usecase

import (
	repo "ginpractice1/internal/infra/repository/shopowner"
	domain "ginpractice1/internal/domain/shopowner"
)

type UpdateProfileRequest struct {
	ID int64 `json:"id"`
	Name string `json:"user"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type ShopOwnerUsecase struct {
	
}

func (u *ShopOwnerUsecase) New() *ShopOwnerUsecase {
	return &ShopOwnerUsecase{}
}

func (u *ShopOwnerUsecase) UpdateMyProfile(request UpdateProfileRequest) bool {
	shopOwner := &domain.Owner{
		ID : int64(request.ID),
		Name: request.Name,
		Email: request.Email,
		Password: request.Password,
	}
	r := repo.New()
	return r.UpdateMyProfile(shopOwner)
}