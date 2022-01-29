// Package configuration - читает из флагов или из переменных окружения,
//заполняет структуру - переменную типа Configuration и возвращает эту структуру и ошибку
//go run hw9.go --port=8080 --dbUrl=postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable --jaegerUrl=http://jaeger:16686 --sentryUrl=http://sentry:9000 --kafkaBroker=kafka:9092 --someAppId=testid --someAppKey=testkey
package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/namsral/flag"
	"log"
	"os"
	//"github.com/joho/godotenv"
)

type Configuration struct {
	port        string `valid:"port"`
	dbUrl       string `valid:"url"`
	jaegerUrl   string `valid:"url"`
	sentryUrl   string `valid:"url"`
	kafkaBroker string `valid:"-"`
	someAppId   string `valid:"-"`
	someAppKey  string `valid:"-"`

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

//lesson8

var set Configuration // Переменная set является структурой типа Configuration
var err error

//Load - читает из флагов или из переменных окружения,
//заполняет структуру - переменную типа Configuration и возвращает эту структуру и ошибку
func Load() (*Configuration, error) {

	//Читаем флаги
	setField1 := flag.String("port", "x", "Enter port")
	setField2 := flag.String("dbUrl", "x", "Enter db_url")
	setField3 := flag.String("jaegerUrl", "x", "Enter jaeger_url")
	setField4 := flag.String("sentryUrl", "x", "Enter sentry_url")
	setField5 := flag.String("kafkaBroker", "x", "Enter kafka_broker")
	setField6 := flag.String("someAppId", "x", "Enter some_app_id")
	setField7 := flag.String("someAppKey", "x", "Enter some_app_key")

	flag.Parse() //Сообщаем библиотеке flag, что необходимо считать флаги
	//Присваеваем значения из флагов
	set.port = *setField1
	set.dbUrl = *setField2
	set.jaegerUrl = *setField3
	set.sentryUrl = *setField4
	set.kafkaBroker = *setField5
	set.someAppId = *setField6
	set.someAppKey = *setField7

	result, err := govalidator.ValidateStruct(set) //Проверяем заполненную структуру валидатором
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	fmt.Println(result) // Выводит true/false

	return &set, err //Вернуть заполненную структуру и ошибку
}

//lesson9
var setJSON Configuration // Переменная set является структурой типа Configuration, для записи конфигурации
//LoadJSON - читает из файла configuration.json,
//заполняет структуру - переменную типа Configuration и возвращает эту структуру и ошибку

func LoadJSON() (*Configuration, error) {
	contentsFileJSON, err := os.ReadFile("configuration.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Print in LoadJSON")
	if err = json.Unmarshal(contentsFileJSON, &setJSON); err != nil {
		log.Fatal(err)
	}
	fmt.Println(setJSON)
	return &setJSON, err //Вернуть заполненную структуру и ошибку
}

//LoadYAML - читает из файла configuration.yaml,
//заполняет структуру - переменную типа Configuration и возвращает эту структуру и ошибку
var setYAML Configuration // Переменная set является структурой типа Configuration, для записи конфигурации
func LoadYAML() (*Configuration, error) {
	fileYAML, err := os.Open("configuration.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.NewDecoder(fileYAML).Decode(&setYAML); err != nil {
		log.Fatal(err)
	}
	return &setYAML, err //Вернуть заполненную структуру и ошибку
}
