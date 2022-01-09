package main

import (
	"encoding/json"
	"fmt"
	"github.com/sqweek/dialog"
	"io/ioutil"
	"os"
)

func main() {
	filename, err := dialog.File().Load()
	if err != nil {
		fmt.Println("Please choose a correct file! ", err)
		os.Exit(1)
	}

	// Функция json.Marshal() возвращает срез байт []byte, который после
	// записи обычно становится не нужен и впоследствии удаляется
	// сборщиком мусора. Если ваша программа ориентирована на массовую
	// обработку и запись JSON, то постоянное выделение и освобождение
	// []byte создает огромную нагрузку на сборщик мусора. Для снижения
	// этой нагрузки лучше применять
	// json.NewEncoder().Encode()
	// который задействует стандартный пакет sync.Pool с целью
	// повторного использования срезов байт []byte для маршализации JSON

	// позволяет прочитать сразу весь файл
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	var parsedData map[string]interface{}
	// cannot read array [{},{}] - but it is normal
	json.Unmarshal([]byte(fileData), &parsedData)

	for key, value := range parsedData {
		fmt.Println("key:", key, "value:", value)
	}
}
