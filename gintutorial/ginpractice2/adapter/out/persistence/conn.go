package persistence

import (
	"database/sql"
	"fmt"
	
	_ "github.com/go-sql-driver/mysql"
	"ginpractice2/adapter/out/persistence/mysql"
)

type MySqlConfig struct {
	Addr     string
	User     string
	Password string
	DB       string
}

func (c *MySqlConfig) NewDB() (*mysql.Queries, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", c.User, c.Password, c.Addr, c.DB)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	queries := mysql.New(db)
	return queries, nil
}