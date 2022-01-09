package main

import (
	"encoding/json"
	"fmt"
	. "leaning/BOOK/200_json/dto"
	"os"
)

func saveToJSON(filename *os.File, key interface{}) {
	encodeJSON := json.NewEncoder(filename) // link to file
	// кодирует данные и сохраняет их в нужный файл
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	myRecord := Record{
		Name:    "Mihalis",
		Surname: "Tsoukalos",
		Tel: []Telephone{
			{Mobile: true, Number: "1234-567"},
			{Mobile: true, Number: "1234-abcd"},
			{Mobile: false, Number: "abcc-567"},
		},
	}
	// os.Stdout - данные выводятся на экран, а не сохраняются в файл
	saveToJSON(os.Stdout, myRecord)
}
