package dbaccess

import (
	"database/sql"
	"fmt"
	
	_ "github.com/go-sql-driver/mysql"
)

type MySqlConfig struct {
	Addr     string
	User     string
	Password string
	DB       string
}

func (c *MySqlConfig) NewDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", c.User, c.Password, c.Addr, c.DB)
	fmt.Printf(dsn)
	return sql.Open("mysql", dsn)
}