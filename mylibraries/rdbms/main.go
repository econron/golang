package main

import (
	"context"
	"database/sql"
	"log"
	"reflect"

	_ "github.com/go-sql-driver/mysql"

	"dbrecipe/tutorial"
)

func run() error {
	ctx := context.Background()

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:13306)/dbname?parseTime=true")
	if err != nil {
		return err
	}

	queries := tutorial.New(db)

	// list all authors
	authors, err := queries.ListAuthros(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	result, err := queries.CreateAutor(ctx, tutorial.CreateAutorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}

	insertedAuthorID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(insertedAuthorID)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthorID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthorID, fetchedAuthor.ID))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}