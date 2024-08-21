package repository

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"ginpractice1/internal/infra/dbaccess"
	shopowner "ginpractice1/internal/domain/shopowner"
)

type ShopOwnerRepository struct {
	query *dbaccess.Queries
}

// dbアクセスする際にここに書いてコネクションプールが使いまわせないのでは？
// だからmain.goに持っていきたい
func New() *ShopOwnerRepository {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:13306)/dbname?parseTime=true")
	if err != nil {
		panic(err)
	}
	queries := dbaccess.New(db)
	return &ShopOwnerRepository{
		query: queries,
	}
}

func (r *ShopOwnerRepository) UpdateMyProfile(owner *shopowner.Owner) bool {
	params := dbaccess.UpdateShopOwnerParams{
		ID: owner.ID,
		Name: owner.Name,
		Email: owner.Email,
		Password: owner.Password,
	}
	// ここcontext.TODOではない気がする
	err := r.query.UpdateShopOwner(context.TODO(), params)
	if err != nil {
		// todo add logger
		return false
	}
	return true
}

func (r *ShopOwnerRepository) CreateAdContent() bool {
	return false
}

func (r *ShopOwnerRepository) UpdateAdContent() bool {
	return true
}