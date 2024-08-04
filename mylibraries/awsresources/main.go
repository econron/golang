package main

import(
	"context"
	"log"
	s3access "awsresources/infrastructure/repository/s3"
)

func main(){
	err := s3access.ListBucketObjects(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}