package actioninfo

import (
	"fmt"
)

// создайте интерфейс DataParser
type DataParser interface {
	Parse(input string) error
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// Попытка распарсить данные
		if err := dp.Parse(data); err != nil {
			// Обработка ошибки парсинга
			fmt.Printf("Ошибка парсинга данных '%s': %v\n", data, err)
			continue // Переход к следующей итерации цикла
		}

		// Вывод информации об активности
		fmt.Println(dp.ActionInfo())
	}
}
