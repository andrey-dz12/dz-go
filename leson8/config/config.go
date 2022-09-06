package configuration

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/kelseyhightower/envconfig"
)

//Struct for conf from example
type Config struct {
	Port        int    `envconfig:"PORT" default:"8080" required:"true"`
	DbUrl       string `envconfig:"DB_URL" default:"postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable" required:"true"`
	JaegerUrl   string `envconfig:"JAEGER_URL" default:"http://jaeger:16686" required:"true"`
	SentryUrl   string `envconfig:"SENTRY_URL" default:"http://sentry:9000" required:"true"`
	KafkaBroker string `envconfig:"KAFKA_BROKER" default:"kafka:9092" required:"true"`
	SomeAppId   string `envconfig:"SOME_APP_ID" default:"testid" required:"true"`
	SomeAppKey  string `envconfig:"SOME_APP_KEY" default:"testkey" required:"true"`
}

//Validate URL in configeration by net/url
func (conf Config) validateConfigURL() error {

	_, err := url.ParseRequestURI(conf.DbUrl)
	if err != nil {
		return errors.New("Incorrect DB URL parameter!")
	}

	_, err = url.ParseRequestURI(conf.JaegerUrl)
	if err != nil {
		return errors.New("Incorrect Jaeger URL parameter!")
	}

	_, err = url.ParseRequestURI(conf.SentryUrl)
	if err != nil {
		return errors.New("Incorrect Sentry URL parameter!")
	}

	return nil
}

//Init configuration
func InitConfig() (Config, error) {
	conf := Config{}
	errProcess := envconfig.Process("", &conf)
	if errProcess != nil {
		fmt.Printf("Can't process the config: %v", errProcess)
		return conf, errProcess
	}

	if errValidate := conf.validateConfigURL(); errValidate != nil {
		fmt.Printf("Incorrect params in config: %v", errValidate)
		return conf, errValidate
	}

	return conf, nil
}
