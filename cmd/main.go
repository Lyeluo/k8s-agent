package main

import (
	"strconv"

	"github.com/spf13/viper"
	_ "k8s.agent/pkg/api"
	"k8s.agent/pkg/config"
	_ "k8s.agent/pkg/registry"
)

func main() {
	r := config.GetWebServer()

	r.Run(":" + strconv.Itoa(viper.GetInt("server.port")))

}
