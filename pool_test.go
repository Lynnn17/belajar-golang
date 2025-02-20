package belajar_golang

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		//default value
		New: func() interface{} {
			return "New"
		},
	}
	group := sync.WaitGroup{}

	pool.Put("Hello")
	pool.Put("World")
	pool.Put("!")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}
	group.Wait()

}
