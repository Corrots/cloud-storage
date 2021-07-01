package main

import (
	"github.com/corrots/cloud-storage/pkg/db"
	"github.com/corrots/cloud-storage/pkg/logging"
	"github.com/corrots/cloud-storage/service"
)

var logger = logging.MustGetLogger("ingress")

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
	logger.Info(val)
	logger.Warn(val)
	logger.Error(val)
	logger.Fatal(val)
}
