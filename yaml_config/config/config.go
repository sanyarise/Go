package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port         string `yaml:"port"`
	Db_url       string `yaml:"db_url"`
	Jaeger_url   string `yaml:"jaeger_url"`
	Sentry_url   string `yaml:"sentry_url"`
	Kafka_broker string `yaml:"kafka_broker"`
	Some_app_id  string `yaml:"some_app_id"`
	Some_app_key string `yaml:"some_app_key"`
}

func Configuration() {
	file, err := os.Open("./config.yaml")
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

	err = yaml.NewDecoder(file).Decode(&Config)
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
