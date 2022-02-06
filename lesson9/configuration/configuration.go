// Package configuration - читает из флагов или из переменных окружения,
//заполняет структуру - переменную типа Configuration и возвращает эту структуру и ошибку
//go run hw9.go --port=8080 --dbUrl=postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable --jaegerUrl=http://jaeger:16686 --sentryUrl=http://sentry:9000 --kafkaBroker=kafka:9092 --someAppId=testid --someAppKey=testkey
package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/namsral/flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
	//"github.com/joho/godotenv"
)

type Configuration struct {
	Port        string `valid:"port" json:"port" yaml:"port"`
	DbUrl       string `valid:"url" json:"dbUrl" yaml:"dbUrl"`
	JaegerUrl   string `valid:"url" json:"jaegerUrl" yaml:"jaegerUrl"`
	SentryUrl   string `valid:"url" json:"sentryUrl" yaml:"sentryUrl"`
	KafkaBroker string `valid:"-" json:"kafkaBroker" yaml:"kafkaBroker"`
	SomeAppId   string `valid:"-" json:"someAppId" yaml:"someAppId"`
	SomeAppKey  string `valid:"-" json:"someAppKey" yaml:"someAppKey"`
}

/* Шаблон:
port: 8080
db_url: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable
jaeger_url: http://jaeger:16686
sentry_url: http://sentry:9000
kafka_broker: kafka:9092
some_app_id: testid
some_app_key: testkey
*/

//lesson8

//LoadFlag - читает из флагов или из переменных окружения,
//заполняет структуру - переменную типа Configuration и возвращает эту структуру и ошибку
func LoadFlag() (*Configuration, error) {
	var set Configuration // Переменная set является структурой типа Configuration
	var err error

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
		fmt.Println("error govalidator.ValidateStruct(set): " + err.Error())
	}
	fmt.Println("govalidator.ValidateStruct(set) result: ", result) // Выводит true/false

	return &set, err //Вернуть заполненную структуру и ошибку
}

//lesson9

//LoadFile - читает из файла configuration.json или configuration.json
//заполняет структуру - переменную типа Configuration и возвращает эту структуру и ошибку
func LoadFile() (*Configuration, error) {
	var setFile Configuration // Переменная set является структурой типа Configuration, для записи конфигурации
	var err error
	var nameConfigFile string
	fmt.Print("Введите название и путь к файлу с конфигурацией json или yaml (configuration/configuration.yaml): ")
	fmt.Scanf("%s\n", &nameConfigFile)

	if strings.Contains(nameConfigFile, ".json") {
		contentsFile, err := os.ReadFile(nameConfigFile)
		if err != nil {
			fmt.Println("Don't read json")
			log.Fatal(err)
		}
		fmt.Print("Print in LoadFile json: ")
		if err = json.Unmarshal(contentsFile, &setFile); err != nil {
			log.Fatal(err)
		}
	}
	//Вариант №2 чтения файла с конфигурацией *.yaml и запись в структру
	if strings.Contains(nameConfigFile, ".yaml") {
		contentsFile, err := ioutil.ReadFile(nameConfigFile)
		if err != nil {
			fmt.Println("Don't read yaml")
			log.Fatal(err)
		}
		fmt.Print("Print in LoadFile yaml 1: ")
		if err = yaml.Unmarshal(contentsFile, &setFile); err != nil {
			log.Fatal(err)
		}
	}
	//Вариант №2 чтения файла с конфигурацией *.yaml и запись в структру
	if strings.Contains(nameConfigFile, ".yaml") {
		contentsFile, err := os.Open(nameConfigFile)
		if err != nil {
			fmt.Println("Don't read yaml ")
			log.Fatal(err)
		}
		fmt.Print("Print in LoadFile yaml: 2")
		if err = yaml.NewDecoder(contentsFile).Decode(&setFile); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(setFile)
	return &setFile, err //Вернуть заполненную структуру и ошибку
}
