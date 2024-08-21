package usecase

import (
	repo "ginpractice1/internal/infra/repository/shopowner"
	domain "ginpractice1/internal/domain/shopowner"
)

// ユースケース層にリクエストの構造体を記述しているのが正直ない
// 何のためのクリーンアーキテクチャ戦略なのか？
type UpdateProfileRequest struct {
	ID int64 `json:"id"`
	Name string `json:"user"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type ShopOwnerUsecase struct {
	
}


// ここでリポジトリを渡せてないのでアウト
// コネクションプールを使いまわせてない
func (u *ShopOwnerUsecase) New() *ShopOwnerUsecase {
	return &ShopOwnerUsecase{}
}

// clean architectureロケット本のweb層の書き方を真似できないか？
// と思って見にいったw
// というのも、usecaseを呼び出す際に、リクエストパラメータのバリデーションなどはそちらに依存している
// usecaseに入れる場合はcqrsパターンの考え方でusecase層を分離して、commandクラスに代入して利用
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