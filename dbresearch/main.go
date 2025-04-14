package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// DSN (Data Source Name) の形式:
	// "user:password@tcp(host:port)/dbname"
	dsn := "myuser:mypassword@tcp(mysql.example.com:3306)/mydatabase"

	db, err := sql.Open("mysql", dsn) // 裏側でgoroutineでXが設定
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 接続確認
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	fmt.Println("Successfully connected to MySQL!")
	name := "test"
	rows, _ := db.Query("SELECT * FROM album WHERE artist = ?", name) // 実際に繋ぎに行ってるのはここ。
	fmt.Println(rows)
}
