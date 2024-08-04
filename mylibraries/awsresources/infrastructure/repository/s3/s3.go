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
	var client *s3.Client

	if(os.Getenv("ENV") == "dev"){
		accessKey := "MYACCESSKEY"
		secretKey := "MYSECRETKEY"
		cred := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")

		endpoint := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: "http://localhost:9000",
			}, nil
		})

		cfg, err := config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(cred), config.WithEndpointResolver(endpoint))
		if err != nil {
			log.Fatalln(err)
		}

		client = s3.NewFromConfig(cfg, func(options *s3.Options) {
			options.UsePathStyle = true
		})
	} else {
		cfg, err = config.LoadDefaultConfig(ctx)
		if err != nil {
			log.Fatal(err)
			return nil, errors.New("Failed to initialize s3 client")
		}
		client = s3.NewFromConfig(cfg)
	}

	return client, nil
}

func ListBucketObjects(ctx context.Context) (*s3.ListObjectsV2Output, error) {
	client, err := s3Client(ctx)
	if err != nil {
		return nil, err
	}
	output, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String("testing-bucket-testtest1234"),
	})
	if err != nil {
		log.Fatalf("Failed to fetch data from s3: %v", err)
		return nil, errors.New("Failed to fetch data from s3.")
	}
	return output, nil
}