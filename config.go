package config

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	port         string
	db_url       string
	jaeger_url   string
	sentry_url   string
	kafka_broker string
	app_id       string
	app_key      string
}

func Conf() {
	var Port = flag.String("port", "8080", "Port number")
	var Db_url = flag.String("db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "Data base address")
	var Jaeger_url = flag.String("jaeger_url", "http://jaeger:16686", "web address jaeger")
	var Sentry_url = flag.String("sentry_url", "http://sentry:9000", "web address sentry")
	var Kafka_broker = flag.String("kafka_broker", "kafka:9092", "kafka id")
	var App_id = flag.String("app_id", "testid", "application identificator")
	var App_key = flag.String("app_key", "testkey", "application key")
	flag.Parse()

	Config := Config{
		port:         *Port,
		db_url:       *Db_url,
		jaeger_url:   *Jaeger_url,
		sentry_url:   *Sentry_url,
		kafka_broker: *Kafka_broker,
		app_id:       *App_id,
		app_key:      *App_key,
	}
	if len(*Port) < 0 || len(*Port) > 4 {
		fmt.Println("Неверное значение номер порта")
		os.Exit(1)
	} else if !strings.Contains(*Db_url, "postgress://") {
		fmt.Println("Неверное значение db_url")
		os.Exit(1)
	} else if !strings.Contains(*Jaeger_url, "http://jaeger") {
		fmt.Println("Неверное значение jaeger_url")
		os.Exit(1)
	} else if !strings.Contains(*Sentry_url, "http://sentry") {
		fmt.Println("Неверное значение sentry_url")
		os.Exit(1)
	} else if !strings.Contains(*Kafka_broker, "kafka:") {
		fmt.Println("Неверное значение kafka_broker")
		os.Exit(1)
	}
	fmt.Println(Config)
}
