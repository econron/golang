package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"fmt"
)

type User struct {
	ID int `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
	Age *int `gorm:"column:age"`
	Active bool `gorm:"column:active"`
	UpdatedAt string `gorm:"column:updated_at"`
}

func main() {
	dsn := "test:password@tcp(localhost:13306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	var user User
	db.First(&user, 2)
	// age := 22
	fmt.Println("1==================")
	db.Model(&user).Updates(User{Name: "new name", Age: nil, Active: true})
	// fmt.Println("2==================")
	// db.Model(&user).Updates(User{Name: "new name", Age: nil, Active: true})
	fmt.Println("2==================")
	db.Model(&user).Select("name", "age").Updates(User{Name: "new-name-2", Age: nil, Active: true})
	fmt.Println("3==================")
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": nil, "active": false})
	// fmt.Println("4==================")
	// db.First(&user, 2)
	// fmt.Printf("user after FIRST: %#v", user)
	// user.Age = nil
	// user.Active = false
	// user.Name = "new-name-2"
	// db.Model(&user).Select("*").Updates(user)

	db.Model(&user).Updates(User{Name: "new name", Age: nil, Active: true}) 


	db.Model(&user).Select("name", "age").Updates(User{Name: "new-name-2", Age: nil, Active: true}) // 2
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": nil, "active": false}) // 3
}

