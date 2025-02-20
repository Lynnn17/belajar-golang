package belajar_golang

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(1 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	channel := time.After(1 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}

	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Dieksekusi Setelah 5 detik")
		group.Done()
	})

	group.Wait()
}
