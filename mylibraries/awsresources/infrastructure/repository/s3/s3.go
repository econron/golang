package s3

import (
	"context"
	"os"
	"log"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func s3Client(ctx context.Context) (*s3.Client, error) {

	var cfg aws.Config
	var err error

	if(os.Getenv("ENV") == "dev"){
		cfg, err = config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("AKID", "SECRET_KEY", "TOKEN")),)
		if err != nil {
			log.Fatal(err)
			return nil, errors.New("Failed to initialize s3 client in development env")
		}
	} else {
		cfg, err = config.LoadDefaultConfig(ctx)
		if err != nil {
			log.Fatal(err)
			return nil, errors.New("Failed to initialize s3 client")
		}
	}
	client := s3.NewFromConfig(cfg)

	return client, nil
}

func ListBucketObjects(ctx context.Context) error {
	client, err := s3Client(ctx)
	if err != nil {
		return err
	}
	output, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String("testing-bucket-testtest1234"),
	})
	if err != nil {
		log.Fatalf("Failed to fetch data from s3.")
		return errors.New("Failed to fetch data from s3.")
	}
	log.Println("first page results:")
	for _, object := range output.Contents {
		log.Printf("key=%s, size=%d", aws.ToString(object.Key), object.Size)
	}
	return nil
}