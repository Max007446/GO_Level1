package yamlP

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"net/url"

	"gopkg.in/yaml.v2"
)

type app struct {
	Port        string `yaml:"port"`
	DbURL       string `yaml:"db_url"`
	JaegerURL   string `yaml:"jaeger_url"`
	SentryURL   string `yaml:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broker"`
	SomeAppID   string `yaml:"some_app_id"`
	SomeAppKey  string `yaml:"some_app_key"`
}

func YamlRead() {
	filename, _ := filepath.Abs("./conf.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Не могу открыть файл: %v", err)
	}

	// Задаем переменную, в которую будем считывать конфиг
	var config app

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Printf("Не могу декодировать yaml-файл в структуру: %v", err)
	}

	jaegerURL, err := url.Parse(config.JaegerURL)
	if err != nil || (jaegerURL.Scheme != "http" && jaegerURL.Scheme != "https") {
		log.Fatalf("Не правильный формат адреса jaeger_url: %v", err)
	}
	sentryURL, err := url.Parse(config.SentryURL)
	if err != nil || (sentryURL.Scheme != "http" && sentryURL.Scheme != "https") {
		log.Fatalf("Не правильный формат адреса sentry_url: %v", err)
	}
	// Теперь в config у нас доступны параметры из файла
	fmt.Printf("Value: %#v\n", config)
}
func yamlP() {
	YamlRead()

}
