package main

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 2)
		wg.Wait() // 阻塞在此
	}()
	wg.Wait() // 阻塞在此
}