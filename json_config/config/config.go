package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	Port         string `json:"port"`
	Db_url       string `json:"db_url"`
	Jaeger_url   string `json:"jaeger_url"`
	Sentry_url   string `json:"sentry_url"`
	Kafka_broker string `json:"kafka_broker"`
	Some_app_id  string `json:"some_app_id"`
	Some_app_key string `json:"some_app_key"`
}

func Configuration() {
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatalf("Не могу открыть файл: %v", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Не могу закрыть файл: %v", err)
		}
	}()
	Config := Config{}

	err = json.NewDecoder(file).Decode(&Config)
	if err != nil {
		log.Printf("Не могу декодировать yaml-файл в структуру: %v", err)
	}
	if !strings.Contains(Config.Db_url, "postgres") {
		log.Fatalln("Неверный Db_URL")
	} else if !strings.Contains(Config.Jaeger_url, "http://jaeger") {
		log.Fatalln("Неверный Jaeger URL")
	} else if !strings.Contains(Config.Sentry_url, "http://sentry") {
		log.Fatalln("Неверный Sentry URL")
	} else if !strings.Contains(Config.Kafka_broker, "kafka") {
		log.Fatalln("Неверное значение поля kafka broker")
	}
	fmt.Println(Config)
}