package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aliceblock/sample/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {

	dataMode := flag.String("mode", "debug", "Data mode {debug|release|dev|prod}")
	flag.Parse()

	switch *dataMode {
	case "dev":
		*dataMode = "debug"
	case "prod":
		*dataMode = "release"
	}

	os.Setenv("DATAMODE", *dataMode)
	config.Init(config.DebugMode)
	if *dataMode == "release" {
		gin.SetMode(gin.ReleaseMode)
		config.Init(config.ReleaseMode)
	}

	// check config
	configInfo := config.Info
	fmt.Println(configInfo.Db.Pass)

	r := gin.Default()

	r.Run()
}
