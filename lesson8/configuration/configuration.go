// Package configuration - читает из флагов или из переменных окружения,
//заполняет структуру - переменную типа Configuration и возвращает эту структуру и ошибку
//go run hw8.go --port=8080 --dbUrl=postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable --jaegerUrl=http://jaeger:16686 --sentryUrl=http://sentry:9000 --someAppId=testid --someAppKey=testkey
package configuration

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/namsral/flag"
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
<<<<<<< HEAD
	fmt.Println(result) // Вывод: true/false
=======
	fmt.Println(result) // Выводит true/false
>>>>>>> 12b1c27a7b898d317fd009129ab3af32c3104a1f

	return &set, err //Вернуть заполненную структуру и ошибку
}
