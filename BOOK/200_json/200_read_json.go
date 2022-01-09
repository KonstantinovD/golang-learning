package main

import (
	"encoding/json"
	"fmt"
	"github.com/sqweek/dialog"
	dtopkg "leaning/BOOK/200_json/dto"
	"os"
)

func loadFromJson(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key) // наполняем данными переданный объект
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

func main() {
	filename, err := dialog.File().Load()
	if err != nil {
		fmt.Println("Please choose a correct file! ", err)
		os.Exit(1)
	}

	var myRecord dtopkg.Record
	err = loadFromJson(filename, &myRecord)
	if err == nil {
		fmt.Println(myRecord)
	} else {
		fmt.Println(err)
	}
}
