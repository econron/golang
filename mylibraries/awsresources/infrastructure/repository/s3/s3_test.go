package s3

import (
	"os"
	"log"
	"testing"
	"github.com/ory/dockertest/v3"
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

    code := m.Run()

    println("after all...")

    os.Exit(code)
}


func TestS3Access(t *testing.T) {
	cases := []struct {
		name string
		want string
	}{
		{
			name: "ファイルを取得できる",
			want: "test_file.txt",
		},
		// {
		// 	name: "ファイルを削除できる",
		// 	want: "",
		// },
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T){
			// ここでconfig注入

			// ここでテスト
		})
	}
}