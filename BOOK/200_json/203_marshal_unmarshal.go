package main

import (
	"encoding/json"
	"fmt"
	. "leaning/BOOK/200_json/dto"
)

func main() {
	myRecord := Record{
		Name:    "Mikhail",
		Surname: "Petrov",
		Tel: []Telephone{
			{Mobile: true, Number: "1234-567"},
			{Mobile: false, Number: "8910-1112"},
		},
	}

	rec, err := json.Marshal(&myRecord)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Marshal:", string(rec))

	var unmarshRecord Record
	err = json.Unmarshal(rec, &unmarshRecord)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Unmarshal:", unmarshRecord)
}
