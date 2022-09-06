package main

import (
	"fmt"
	"leson8/config/config"
	"os"
)

func main() {
	conf, err := config.InitConfig()
	if err != nil {
		fmt.Printf("Error load configuration\n")
		os.Exit(1)
	}

	fmt.Printf("Configuration from ENV: %+v\n\n", conf)
	fmt.Printf("Configuration load OK\n")
}

