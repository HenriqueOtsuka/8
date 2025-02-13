package main

import (
	"github.com/HenriqueOtsuka/8/configs/config"
)

func main() {
	config, _ := config.LoadConfig(".env")
	println(config.AppPort)
}
