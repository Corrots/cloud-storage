package main

import (
	"fmt"

	"github.com/corrots/cloud-storage/db"
	"github.com/corrots/cloud-storage/service"
)

func main() {
	if err := db.InitRedis(":6379", 0, ""); err != nil {
		panic(err)
	}
	cacheSvc := service.NewCacheService()
	if err := cacheSvc.Set("k1", "100"); err != nil {
		panic(err)
	}

	var val string
	if err := cacheSvc.GetI("k1", &val); err != nil {
		panic(err)
	}
	fmt.Println(val)
}
