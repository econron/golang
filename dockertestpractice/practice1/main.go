package main

import (
	sqlc "practice1/sqlc"
	dbaccess "practice1/dbaccess"
	_ "github.com/go-sql-driver/mysql"
	"context"
	"database/sql"

)

func main() {
	config := &dbaccess.MySqlConfig{
		Addr: "127.0.0.1:13306",
		User: "user",
		Password: "password",
		DB: "dbname",
	}
	db, err := config.NewDB()
	if err != nil {
		panic(err)
	}
	TransferFunds(db, "100", "200")
}

func TransferFunds(db *sql.DB, b1 string, b2 string) error {
	q := sqlc.New(db)

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()
	qtx := q.WithTx(tx)

	arg1 := sqlc.UpdateUserBalanceParams{
		Balance: b1,
		ID: 1,
	}
	arg2 := sqlc.UpdateUserBalanceParams{
		Balance: b2,
		ID: 2,
	}
	arg3 := sqlc.RecordTransactionParams{
		SenderID: 2,
		ReceiverID: 1,
		Amount: "100",
	}
	err1 := qtx.UpdateUserBalance(context.TODO(), arg1)
	if err1 != nil {
		return err
	}
	err2 := qtx.UpdateUserBalance(context.TODO(), arg2)
	if err2 != nil {
		return err
	}
	err3 := qtx.RecordTransaction(context.TODO(), arg3)
	if err3 != nil {
		return err
	}
	return tx.Commit()
}