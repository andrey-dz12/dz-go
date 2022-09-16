package configuration

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/kelseyhightower/envconfig"
	"conf.yaml"
)

// Struct for conf from example
type Config struct {
	Port        int    `envconfig:"PORT" default:"8080" required:"true" json:"port" yaml:"port"`
	DbUrl       string `envconfig:"DB_URL" default:"postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable" required:"true" json:"db_url" yaml:"db_url"`
	JaegerUrl   string `envconfig:"JAEGER_URL" default:"http://jaeger:16686" required:"true" json:"jaeger_url" yaml:"jaeger_url"`
	SentryUrl   string `envconfig:"SENTRY_URL" default:"http://sentry:9000" required:"true" json:"sentry_url" yaml:"sentry_url"`
	KafkaBroker string `envconfig:"KAFKA_BROKER" default:"kafka:9092" required:"true" json:"kafka_broker" yaml:"kafka_broker"`
	SomeAppId   string `envconfig:"SOME_APP_ID" default:"111" required:"true" json:"some_app_id" yaml:"some_app_id"`
	SomeAppKey  string `envconfig:"SOME_APP_KEY" default:"dev" required:"true" json:"some_app_key" yaml:"some_app_key"`
}

//Validate URL in configeration by net/url
func (conf *Config) validateConfigURL() error {

	_, err := url.ParseRequestURI(conf.DbUrl)
	if err != nil {
		return errors.New("incorrect db url parameter")
	}

	_, err = url.ParseRequestURI(conf.JaegerUrl)
	if err != nil {
		return errors.New("incorrect jaeger url parameter")
	}

	_, err = url.ParseRequestURI(conf.SentryUrl)
	if err != nil {
		return errors.New("incorrect sentry url parameter")
	}

	return nil
}

// Print config
func (conf Config) PrintConfig(source string) error {
	_, err := fmt.Printf("Current configuration by %s is: %+v\n\n", source, conf)
	return err
}

// Init configuration for enviroment, JSON and YAML
func InitConfig(source string, configPath string) (Config, error) {
	conf := Config{}
	switch source {
	case "env":
		return conf, loadConfigEnv(&conf)
	case "json":
		return conf, loadConfigJson(&conf, configPath)
	case "yaml":
		return conf, loadConfigYaml(&conf, configPath)
	default:
		return conf, errors.New("unknown configuration")
	}
	return conf, nil
}

// Load config from enviroment
func loadConfigEnv(conf *Config) error {
	errProcess := envconfig.Process("", conf)
	if errProcess != nil {
		fmt.Printf("Can't process the config: %v", errProcess)
		return errProcess
	}

	if errValidate := conf.validateConfigURL(); errValidate != nil {
		fmt.Printf("Incorrect params in config ENV: %v", errValidate)
		return errValidate
	}

	return nil
}

// Load config from JSON file
func loadConfigJson(conf *Config, configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Error on close config file: %v", err)
		}
	}()

	err = json.NewDecoder(file).Decode(conf)
	if err != nil {
		return err
	}

	if errValidate := conf.validateConfigURL(); errValidate != nil {
		fmt.Printf("Incorrect params in config JSON: %v", errValidate)
		return errValidate
	}

	return nil
}

// Load config from YAML file
func loadConfigYaml(conf *Config, configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Error on close config file: %v", err)
		}
	}()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(byteValue), conf)
	if err != nil {
		return err
	}

	if errValidate := conf.validateConfigURL(); errValidate != nil {
		fmt.Printf("Incorrect params in config YAML: %v", errValidate)
		return errValidate
	}

	return nil
}