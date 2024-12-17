package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
)

var(
	kGoBuildSuccess          = tag.MustNewKey("go-playground/frontend/go_build_success")
	mGoBuildLatency          = stats.Float64("go-playground/frontend/go_build_latency", "", stats.UnitMilliseconds)
)

func main() {
	tmpDir, err := os.MkdirTemp("", "sandbox")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpDir)
	fmt.Printf(tmpDir)

	start := time.Now()
	defer func() {
		status := "success"
		if err != nil {
			status = "error"
		}
		stats.RecordWithTags(context.TODO(), []tag.Mutator{tag.Upsert(kGoBuildSuccess, status)}, mGoBuildLatency.M(float64(time.Since(start))/float64(time.Millisecond)))
	}()
}