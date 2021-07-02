package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var mutex sync.Mutex

func main() {
	file, err := os.Create("./tmp/a.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rand.Seed(time.Now().UnixNano())

	done := make(chan struct{}, 1)

	go func() {
		mutex.Lock()
		// file input
		file.Write([]byte(time.Now().Format("2006-01-02 15:04:05\n")))
		mutex.Unlock()
	}()

	timer := time.NewTimer(2 * time.Second)
	go func() {
		for {
			select {
			case <-timer.C:
				done <- struct{}{}
				return
			default:
				mutex.Lock()
				file.WriteString(time.Now().Format(time.RFC3339Nano) + "\n")
				mutex.Unlock()
			}
		}
	}()

	<-done
	fmt.Println("done")
}
