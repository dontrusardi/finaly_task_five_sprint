package actioninfo

import (
	"log"
	"fmt"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, element := range dataset {
		values := dp.Parse(element)
		if values != nil {
			log.Printf("error parsing: %v, string: %s", values, element)
			continue
		}
		actionInfo, err := dp.ActionInfo()
		if err != nil {
			log.Printf("information error: %v, string: %s", err, element)
			continue
		}
		fmt.Println(actionInfo)
	}
}
