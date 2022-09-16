package main

import (
	"flag"
	"fmt"
	"leson8/config/config"
	"os"
)

var jsonConfigPath string
var yamlCongigPath string

func main() {

	flag.StringVar(&jsonConfigPath, "json_config_path", "./configuration/conf.json", "Path for JSON config file")
	flag.StringVar(&yamlCongigPath, "yaml_config_path", "./configuration/conf.yaml", "Path for YAML config file")

	conf, err := config.InitConfig("env", "")
	if err != nil {
		fmt.Printf("Error load configuration: %s\n", err)
		os.Exit(1)
	}
	conf.PrintConfig("ENV")

	conf, err = config.InitConfig("json", jsonConfigPath)
	if err != nil {
		fmt.Printf("Error load configuration: %s\n", err)
		os.Exit(1)
	}
	conf.PrintConfig("JSON")

	conf, err = config.InitConfig("yaml", yamlCongigPath)
	if err != nil {
		fmt.Printf("Error load configuration: %s\n", err)
		os.Exit(1)
	}
	conf.PrintConfig("YAML")

	fmt.Printf("Configuration load OK\n")
}

