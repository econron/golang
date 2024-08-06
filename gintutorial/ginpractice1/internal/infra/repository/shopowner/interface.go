package repository

type ShopOwnerRepositoryInterface interface {
	UpdateMyProfile() bool
	CreateAdContent() bool
	UpdateAdContent() bool
}