package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Ardi")
	pool.Put("Mardiansyah")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			time.Sleep(1 * time.Second)
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Selesai")
}
