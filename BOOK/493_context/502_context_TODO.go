package main

import (
	"context"
	"fmt"
)

// создадим контекст, в котором вместо функции context.Background()
// будет использоваться context.TODO
// Обе эти функции возвращают ненулевой пустой контекст (= одно и то же),
// однако назначение у них разное

type aKey string

// функция извлекает значение из контекста и проверяет,
// существует ли такое значение.
func searchKey(ctx context.Context, k aKey) {
	// Value returns the value associated with this context for key
	v := ctx.Value(k)
	if v != nil {
		fmt.Println("found value:", v)
		return
	} else {
		fmt.Println("key not found:", k)
	}
}

func main() {
	myKey := aKey("mySecretValue")
	// --- read WithValue doc
	// --- context.WithValue() предоставляет способ связывания значения
	// с объектом Context.
	// --- Контексты не должны храниться в структурах — их следует
	// передавать в функции в виде отдельных аргументов. Рекомендуется
	// передавать контекст в качестве первого аргумента функции
	// -- The provided key must be comparable and should not be of type
	// string or any other built-in type to avoid collisions between
	// packages using context
	ctx := context.WithValue(
		context.Background(), myKey, "mySecretValue")
	searchKey(ctx, myKey)
	searchKey(ctx, aKey("notThere"))

	// --- В данном случае мы заявляем, что, несмотря на то что мы
	// намерены использовать контекст операции, мы все еще не уверены в
	// этом и потому используем функцию context.TODO().
	// --- Эта функция распознается средствами статического анализа, что
	// позволяет им определить, правильно ли распространяются контексты
	emptyCtx := context.TODO()
	searchKey(emptyCtx, aKey("mySecretValue"))
	// --- Запомните: никогда не передавайте нулевой контекст —
	// используйте функцию context.TODO()
	// для создания подходящего контекста
}
