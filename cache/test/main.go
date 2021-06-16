package main

import (
	"fmt"

	redis2 "github.com/corrots/cloud-storage/cache/redis"
)

func main() {
	c := redis2.NewConn()
	defer c.Close()
	res, err := c.Do("PING")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", res)
}
