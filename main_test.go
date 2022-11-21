package thingy

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestThingy(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	worker := New()
	go worker.Run(ctx)

	worker.Upsert("test", "test")
	worker.Upsert("bingo", "bongo")
	worker.Upsert("test", "whatsup")
	worker.Delete("sdf")
	worker.Delete("test")

	fmt.Println(worker.Done())
}
