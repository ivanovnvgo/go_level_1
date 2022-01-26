//Урок 8. Основы конфигурирования приложений
//Разработайте пакет для чтения конфигурации типичного веб-приложения через флаги или переменные окружения.
//Пример конфигурации можно посмотреть здесь. По желанию вы можете задать другие имена полям, сгруппировать их или добавить собственные поля.
//Помимо чтения конфигурации приложение также должно валидировать её - например, все URL’ы должны соответствовать ожидаемым форматам.
//Работу с конфигурацией необходимо вынести в отдельный пакет (не в пакет main).
package main

import (
	"fmt"
	"go_level_1/go_level_1/lesson8/configuration"
	"os"
)

func main() {

	config, err := configuration.Load()
	if err != nil {
		fmt.Println("config could be loaded: %w", err)
		os.Exit(1)
	}
	fmt.Printf("%#v\n", config)
}
