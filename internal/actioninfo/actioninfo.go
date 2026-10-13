package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(dataset string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for i := range dataset {

		if err := dp.Parse(dataset[i]); err != nil {
			log.Printf("ошибка парсинга: %v", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка получения информации: %v", err)
			continue
		}

		fmt.Println(info)
	}
}
