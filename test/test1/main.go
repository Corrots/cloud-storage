package main

import (
	"fmt"
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

	done := make(chan struct{}, 1)

	//for {
	//	select {
	//	case <-timer.C:
	//		return
	//	case <-done:
	//		return
	//	default:
	//		file.WriteString("main goroutine: \n")
	//	}
	//}

	go func() {
		//mutex.Lock()
		// file input
		file.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))

		//mutex.Unlock()
		done <- struct{}{}
		fmt.Println("done")
	}()

	go func() {

	}()

	<-done
}
