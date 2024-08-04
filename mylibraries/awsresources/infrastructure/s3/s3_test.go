package s3

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/ory/dockertest/v3"
	dc "github.com/ory/dockertest/v3/docker"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func TestMain(m *testing.M) {
    println("before all...")

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	options := &dockertest.RunOptions{
		Repository: "minio/minio",
		Tag:        "latest",
		Cmd:        []string{"server", "/data"},
		PortBindings: map[dc.Port][]dc.PortBinding{
			"9000/tcp": []dc.PortBinding{{HostPort: "9000"}},
		},
		Env: []string{"MINIO_ACCESS_KEY=MYACCESSKEY", "MINIO_SECRET_KEY=MYSECRETKEY"},
	}

	resource, err := pool.RunWithOptions(options)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	endpoint := fmt.Sprintf("localhost:%s", resource.GetPort("9000/tcp"))
	url := fmt.Sprintf("http://%s/minio/health/live", endpoint)

	fmt.Println(url)

	if err := pool.Retry(func() error {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("status code not OK")
		}
		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	err = setUpBucket(endpoint)
	if err != nil {
		log.Fatal("Failed to setup minio bucket.")
	}

    code := m.Run()

	defer func() {
		if err := pool.Purge(resource); err != nil {
		  log.Fatalf("Could not purge resource: %s", err)
		}
	}()

    println("after all...")

    os.Exit(code)
}

func setUpBucket(endpoint string) error {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4("MYACCESSKEY", "MYSECRETKEY", ""),
		Secure: false,
	})

	if err != nil {
		log.Fatal("Failed to access to minio.")
		log.Fatalln(err)
		panic(err)
	}

	// ファイルを作成
    filePath := "can_get_test_file.txt"
    file, err := os.Create(filePath)
    if err != nil {
        log.Fatalln(err)
    }

    // ファイルにデータを書き込み
    _, err = file.WriteString("This is a test file for MinIO upload.")
    if err != nil {
        log.Fatalln(err)
    }
    file.Close()

    // アップロードのためにファイルを再度開く
    file, err = os.Open(filePath)
    if err != nil {
        log.Fatalln(err)
    }
    defer file.Close()

    // ファイル情報を取得
    fileInfo, err := file.Stat()
    if err != nil {
        log.Fatalln(err)
    }

    // バケット作成
    bucketName := "testing-bucket-testtest1234"
	exists, err := minioClient.BucketExists(context.Background(), bucketName)
    if err != nil {
        log.Fatalln(err)
    }
    if !exists {
        err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
        if err != nil {
            log.Fatalln(err)
        }
        log.Println("バケットを作成しました:", bucketName)
    }

	// ファイルをアップロード
    _, err = minioClient.PutObject(context.Background(), bucketName, filePath, file, fileInfo.Size(), minio.PutObjectOptions{
        ContentType: "application/octet-stream",
    })
    if err != nil {
        log.Fatalln(err)
		return err
    }
	log.Println("ファイルをアップロードしました:", filePath)
	return nil
}


func TestS3Access(t *testing.T) {
	cases := []struct {
		name string
		want string
	}{
		{
			name: "ファイルを取得できる",
			want: "can_get_test_file.txt",
		},
		// {
		// 	name: "ファイルを削除できる",
		// 	want: "",
		// },
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T){
			output, err := ListBucketObjects(context.TODO())
			if err != nil {
				t.Error("error")
			}
			for _, object := range output.Contents {
				got := aws.ToString(object.Key)
				if got != tt.want {
					t.Errorf("want = %s, but got = %s", tt.want, got)
				}
			}
		})
	}
}