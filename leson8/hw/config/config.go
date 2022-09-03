import (
	"fmt"
	"os"
	"regexp"
)

type SrvConfigs struct {

}

func (c *SrvConfigs) DefaultConfigs() {

	if c.Port == "" {
	matchPort, _ :=regexp.MatchString('[0-9a-z_]/+',c.port)
	if c.Port == ""|| !matchPort {
		c.Port = "8080"
	} 
} 
	if c.Db_url == "" {
	matchDb_url, _ :=regexp.MatchString('[0-9a-z_]/+',c.port)
	if c.Db_url == ""|| !matchPort {
		c.Db_url = "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"
	} 
} 
if c.Jaeger_url == "" {
	matchJaeger_url, _ :=regexp.MatchString('[0-9a-z_]/+',c.port)
	if c.Jaeger_url == ""|| !matchPort {
		c.Jaeger_url = "http://jaeger:16686"
	} 
} 
if c.Sentry_url == "" {
	matchSentry_url, _ :=regexp.MatchString('[0-9a-z_]/+',c.port)
	if c.Sentry_url == ""|| !matchPort {
		c.Sentry_url = "http://sentry:9000"
	} 
} 
if c.Kafka_broker == "" {
	matchKafka_broker, _ :=regexp.MatchString('[0-9a-z_]/+',c.port)
	if c.Kafka_broker == ""|| !matchPort {
		c.Kafka_broker = "9092"
	} 
} 
if c.Some_app_id == "" {
	matchSome_app_id, _ :=regexp.MatchString('[0-9a-z_]/+',c.port)
	if c.Some_app_id == ""|| !matchPort {
		c.Some_app_id = "testid"
	} 
} 
if c.Some_app_key == "" {
	matchSome_app_key, _ :=regexp.MatchString('[0-9a-z_]/+',c.port)
	if c.Some_app_key == ""|| !matchPort {
		c.Some_app_key = "testid"
	} 
} 

some_app_key: testkey
	
}