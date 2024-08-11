package main

import (
	"context"
	"os"

	"golang.org/x/exp/slog"

	"txdi/infra/mysql"
	"txdi/infra/mysql/repository"
	"txdi/usecase"
)

func main() {
	mysqlDB, err := mysql.NewDB(&mysql.Config{
		Addr:     os.Getenv("MYSQL_ADDR"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DB:       os.Getenv("MYSQL_DATABASE"),
	})
	
	if err != nil {
		slog.Error(err.Error())
		return
	}

	txManager := mysql.NewTxManager(mysqlDB)
	repository := repository.NewUserRepository()
	interactor := usecase.NewUserInteractor(txManager, repository)

	ctx := context.Background()
	userId := "user_id"
	if _, err := interactor.GetUser(ctx, userId); err != nil {
		slog.Error(err.Error())
		return
	}

	if err := interactor.UpdateName(ctx, userId, "test"); err != nil {
		slog.Error(err.Error())
		return
	}
}