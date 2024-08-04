package main

import(
	"context"
	"log"
	// "fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main(){
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	client := s3.NewFromConfig(cfg)

	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("testing-bucket-testtest1234"),
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("first page results:")
	for _, object := range output.Contents {
		log.Printf("key=%s, size=%d", aws.ToString(object.Key), object.Size)
	}
}