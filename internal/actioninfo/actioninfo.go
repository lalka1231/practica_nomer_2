package actioninfo

import "fmt"

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for i := range dataset {
		err := dp.Parse(dataset[i])
		if err != nil {
			fmt.Printf("Ошибка парсинга: %v\n", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("Ошибка формирования отчета: %v\n", err)
			continue
		}
		fmt.Println(info)
		fmt.Println()
	}
}
