package repository

import (
	"context"
	"database/sql"
	"ginpractice1/internal/infra/dbaccess"
	shopowner "ginpractice1/internal/domain/shopowner"
)

type ShopOwnerRepository struct {
	query *dbaccess.Queries
}

func (*ShopOwnerRepository) New() *ShopOwnerRepository {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:13306)/dbname?parseTime=true")
	if err != nil {
		panic(err)
	}
	queries := dbaccess.New(db)
	return &ShopOwnerRepository{
		query: queries,
	}
}

func (r *ShopOwnerRepository) UpdateMyProfile(ctx context.Context, owner shopowner.Owner) bool {
	params := dbaccess.UpdateShopOwnerParams{
		ID: owner.ID,
		Name: owner.Name,
		Email: owner.Email,
		Password: owner.Password,
	}
	r.query.UpdateShopOwner(ctx, params)
	return false
}

func (r *ShopOwnerRepository) CreateAdContent() bool {
	return false
}

func (r *ShopOwnerRepository) UpdateAdContent() bool {
	return true
}