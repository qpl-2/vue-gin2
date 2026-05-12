package main

import (
	"vuegin/config"
	"vuegin/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()

	r.Run()
}
