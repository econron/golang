package repository

import (
	"context"

	"txdi/domain/entity"
	"txdi/domain/transaction"
)

type UserRepository interface {
	SelectByPK(ctx context.Context, tx transaction.ROTx, userID string) (*entity.User, error)
	Update(ctx context.Context, tx transaction.RWTx, user *entity.User) error
}