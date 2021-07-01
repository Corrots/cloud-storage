package main

import (
	"os"
	"sync"
	"time"
)

var mutex sync.Mutex

func main() {
	file, err := os.Create("./tmp/a.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("main goroutine: ")

	done := make(chan struct{}, 1)

	go func() {
		mutex.Lock()
		// file input
		file.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))

		mutex.Unlock()
		done <- struct{}{}
	}()

	<-done
}
