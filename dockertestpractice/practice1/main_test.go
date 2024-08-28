package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"

	sqlc "practice1/sqlc"
)

var db *sql.DB

func TestMain(m *testing.M) {
	var db *sql.DB
	var err error
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err := pool.Run("mysql", "8.0.29", []string{"MYSQL_ROOT_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql?parseTime=true", resource.GetPort("13306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	defer func(){
		if err = pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}()

	m.Run()
	os.Exit(resource.Container.State.ExitCode)
}

func TestTransferFunc(t *testing.T) {
	cases := []struct{
		name string
		b1 string
		b2 string
		transactionId int32
		want int32
	}{
		{
			name: "正常系1",
			b1: "100",
			b2: "200",
			transactionId: 1,
			want: 1,
		},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T){
			err := TransferFunds(db, tt.b1, tt.b2)
			if err != nil {
				t.Error(err.Error())
			}
			got := getTransactionId(db, tt.transactionId)
			if got.ID != tt.want {
				t.Errorf("got %d, want %d", got.ID, tt.want)
			}
		})
	}
}

func getTransactionId(db *sql.DB, id int32) sqlc.Transaction {
	q := sqlc.New(db)
	ret, _ := q.GetTransaction(context.TODO(), id)
	return ret
}