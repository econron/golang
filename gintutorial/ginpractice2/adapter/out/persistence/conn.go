package persistence

import (
	"database/sql"
	
	_ "github.com/go-sql-driver/mysql"
	"ginpractice2/adapter/out/persistence/mysql"
)

type MySqlConfig struct {
}

func (c *MySqlConfig) NewDB() (*mysql.Queries, error) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:13306)/dbname?parseTime=true")
	if err != nil {
		panic(err)
	}
	queries := mysql.New(db)
	return queries, nil
}