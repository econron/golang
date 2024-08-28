package practice2

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"

)

func TestHoge(t *testing.T) {
	var db *sql.DB
var err error
pool, err := dockertest.NewPool("")
pool.MaxWait = time.Minute * 2
if err != nil {
    log.Fatalf("Could not construct pool: %s", err)
}

err = pool.Client.Ping()
if err != nil {
    log.Fatalf("Could not connect to Docker: %s", err)
}

resource, err := pool.Run("mysql", "8.0.29", []string{"MYSQL_ROOT_PASSWORD=secret"})
if err != nil {
    log.Fatalf("Could not start resource: %s", err)
}

if err = pool.Retry(func() error {
    var err error
    db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql?parseTime=true", resource.GetPort("3306/tcp")))
    if err != nil {
        return err
    }
    return db.Ping()
}); err != nil {
    log.Fatalf("Could not connect to docker: %s", err)
}

// When you're done, kill and remove the container
if err = pool.Purge(resource); err != nil {
    log.Fatalf("Could not purge resource: %s", err)
}
}