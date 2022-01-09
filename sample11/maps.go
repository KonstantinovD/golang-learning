package main

import (
	"fmt"
)

func main() {
	// --- map представляет ссылку на хеш-таблицу - структуру данных,
	// где каждый элемент представляет пару "ключ-значение". При этом
	// каждый элемент имеет уникальный ключ, по которому можно получить
	// значение элемента.
	// Отображение определяется как объект типа map[K]V,
	// где К представляет тип ключа, а V - тип значения.
	// --- !!! Тип ключа K должен поддерживать операцию сравнения ==,
	//чтобы отображение могло сопоставить значение с ключом из хеш-таблицы

	// --- The zero value of a map is nil. A nil map has no keys,
	// nor keys can be added.
	// --- The make() func returns a map of the given type.
	// ex: make(map[string]int) - keys - string, values - int

	var people = map[string]int{ // keys - string, values - int
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}
	fmt.Println(people) // map[Tom:1 Bob:2 Sam:4 Alice:8]

	fmt.Println(people["Alice"]) // 8
	fmt.Println(people["Bob"])   // 2
	people["Bob"] = 32
	fmt.Println(people["Bob"]) // 32
	fmt.Println()

	// --- Для проверки наличия элемента по определенному ключу
	// можно применять выражение if:
	if val, ok := people["Tom"]; ok {
		fmt.Println(val)
	}
	// EQUIVALENT TO
	val, ok := people["Tom"]
	if ok {
		fmt.Println(val)
	}
	fmt.Println()

	// --- Для перебора элементов применяется цикл for:
	for key, value := range people {
		fmt.Println(key, value)
	}
	fmt.Println()

	// --- Функция make создает пустую хеш-таблицу:
	names := make(map[string]int)
	names["Kate"] = 128 // Добавление нового элемента
	names["Tom"] = 256
	fmt.Println("Filled map:", names) // map[Kate:128 Tom:256]

	// -- Для удаления применяется встроенная функция
	// delete(map, key),
	delete(names, "Bob")                      // nothing
	fmt.Println("After deleting Bob:", names) //map[Kate:128 Tom:256]
	delete(names, "Kate")
	fmt.Println("After deleting Kate:", names) // map[Tom:256]

}
