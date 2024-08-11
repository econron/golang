package usecase

import (
	"context"

	"txdi/domain/entity"
	"txdi/domain/repository"
	"txdi/domain/transaction"
)

type UserInteractor interface {
	GetUser(ctx context.Context, userID string) (*entity.User, error)
	UpdateName(ctx context.Context, userID, name string) error
}

type userInteractor struct {
	txManager      transaction.TxManager
	userRepository repository.UserRepository
}

func NewUserInteractor(
	txManager transaction.TxManager,
	userRepository repository.UserRepository,
) UserInteractor {
	return &userInteractor{
		txManager,
		userRepository,
	}
}

func (i *userInteractor) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	var user *entity.User
	if err := i.txManager.ReadOnlyTransaction(ctx, func(ctx context.Context, tx transaction.ROTx) error {
		var err error
		user, err = i.userRepository.SelectByPK(ctx, tx, userID)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return user, nil
}

func (i *userInteractor) UpdateName(ctx context.Context, userID, name string) error {
	if err := i.txManager.ReadWriteTransaction(ctx, func(ctx context.Context, tx transaction.RWTx) error {
		user, err := i.userRepository.SelectByPK(ctx, tx, userID)
		if err != nil {
			return err
		}
		user.Name = name
		if err := i.userRepository.Update(ctx, tx, user); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}