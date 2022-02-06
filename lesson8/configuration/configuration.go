// Package configuration - читает из флагов или из переменных окружения,
//заполняет структуру - переменную типа Configuration и возвращает эту структуру и ошибку
//go run hw8.go --port=8080 --dbUrl=postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable --jaegerUrl=http://jaeger:16686 --sentryUrl=http://sentry:9000 --someAppId=testid --someAppKey=testkey
package configuration

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/namsral/flag"
	"log"
	"net/url"
)

type Configuration struct {
	Port        string `valid:"port"`
	DbUrl       string `valid:"url"`
	JaegerUrl   string `valid:"url"`
	SentryUrl   string `valid:"url"`
	KafkaBroker string `valid:"-"`
	SomeAppId   string `valid:"-"`
	SomeAppKey  string `valid:"-"`
	/* Шаблон:
	port: 8080
	db_url: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable
	jaeger_url: http://jaeger:16686
	sentry_url: http://sentry:9000
	kafka_broker: kafka:9092
	some_app_id: testid
	some_app_key: testkey
	*/
}

var set Configuration // Переменная set является структурой типа Configuration
var err error

//Load - читает из флагов или из переменных окружения,
//заполняет структуру - переменную типа Configuration и возвращает эту структуру и ошибку
func Load() (*Configuration, error) {

	//Читаем флаги и присваиваем значения полям структуры set у которой тип Configeration
	flag.StringVar(&set.Port, "port", "x", "Enter port")
	flag.StringVar(&set.DbUrl, "dbUrl", "x", "Enter db_url")
	flag.StringVar(&set.JaegerUrl, "jaegerUrl", "x", "Enter jaeger_url")
	flag.StringVar(&set.SentryUrl, "sentryUrl", "x", "Enter sentry_url")
	flag.StringVar(&set.KafkaBroker, "kafkaBroker", "x", "Enter kafka_broker")
	flag.StringVar(&set.SomeAppId, "someAppId", "x", "Enter some_app_id")
	flag.StringVar(&set.SomeAppKey, "someAppKey", "x", "Enter some_app_key")

	flag.Parse() //Сообщаем библиотеке flag, что необходимо считать флаги

	//Mоя валидация 1 url
	u, err := url.Parse(set.DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Scheme: ", u.Scheme)
	fmt.Println("Host: ", u.Host)
	fmt.Println("Path: ", u.Path)
	fmt.Println("RawQuery: ", u.RawQuery)

	//Mоя валидация 2 url
	const dbUrl = "/petstore?sslmode=disable"
	u, err = url.Parse(set.DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	if u.RequestURI() != dbUrl {
		log.Fatal("u.RequestURL", err)
	}

	fmt.Println("End of the user validation") // Delete

	//Валидация при помощи установленной сторонней библиотеки
	result, err := govalidator.ValidateStruct(set) //Проверяем заполненную структуру валидатором
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	fmt.Println(result) // Вывод: true/false

	return &set, err //Вернуть заполненную структуру и ошибку
}
