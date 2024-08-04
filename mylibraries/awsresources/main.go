package main

import(
	"context"
	"log"
	s3access "awsresources/infrastructure/repository/s3"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func main(){
	output, err := s3access.ListBucketObjects(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("first page results:")
	for _, object := range output.Contents {
		log.Printf("key=%s, size=%d", aws.ToString(object.Key), object.Size)
	}
}