package main

// Библиотека reflewalk позволяет перебирать сложные значения в Go,
// используя рефлексию, — подобно тому как мы перебираем файлы
// в файловой системе

import (
	"fmt"
	"github.com/mitchellh/reflectwalk"
	"reflect"
)

type Values struct {
	Extra map[string]string
}

type WalkMap struct {
	MapVal reflect.Value
	Keys   map[string]bool
	Values map[string]bool
}

// Map -> Интерфейс, определеннный в reflewalk, требует реализации
// функции Map(), которая используется для поиска в хеш-таблицах
func (t *WalkMap) Map(m reflect.Value) error {
	t.MapVal = m
	return nil
}

func (t *WalkMap) MapElem(m, k, v reflect.Value) error {
	if t.Keys == nil {
		t.Keys = make(map[string]bool)
		t.Values = make(map[string]bool)
	}

	t.Keys[k.Interface().(string)] = true // see type assertions
	t.Keys[v.Interface().(string)] = true
	return nil
}

func main() {
	walkedmap := new(WalkMap)

	type S struct {
		Map map[string]string
	}

	// определяем новую переменную data, которая содержит хеш-таблицу,
	data := &S{
		Map: map[string]string{
			"V1": "v1v",
			"V2": "v2v",
			"V3": "v3v",
			"V4": "v4v",
		},
	}
	// и вызываем функцию reflewalk.Walk(), чтобы ее исследовать
	err := reflectwalk.Walk(data, walkedmap)
	if err != nil {
		fmt.Println(err)
		return
	}

	// рефлексия для вывода на экран содержимого поля MapVal
	// структуры WalkMap
	r := walkedmap.MapVal // поле структуры
	fmt.Println("MapVal:", r)
	rType := r.Type()
	fmt.Printf("Type of r: %s\n", rType)
	// MapKeys() возвращает срез reflect.Values, каждое значение
	// которого содержит один из ключей хеш-таблицы
	for _, key := range r.MapKeys() {
		// MapIndex() позволяет вывести на экран значение ключа
		fmt.Println("key:", key, "value:", r.MapIndex(key))
	}
	// --- Методы MapKeys() и MapIndex() работают только с типом reflect.Map
	// и позволяют перебирать все значения хеш-таблицы;
	// --- Однако последовательность возвращаемых элементов
	// хеш-таблицы будет случайной
}
