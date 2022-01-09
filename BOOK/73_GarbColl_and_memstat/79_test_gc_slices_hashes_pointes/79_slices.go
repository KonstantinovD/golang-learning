package main

import (
	"fmt"
	"runtime"
	"time"
)

type data struct {
	i, j int
}

func main() {
	start := time.Now()

	var N = 40000000
	var structure []data
	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
	}

	runtime.GC()
	_ = structure[0]
	// оператор (_ = structure[0]) используется для предотвращения
	// преждевременной очистки сборщиком мусора переменной structure

	elapsed := time.Since(start) // Execution took 2.3756601s
	fmt.Printf("Execution slices took %s", elapsed)

	// --- хеш-таблицы замедляют сборщик мусора Go,
	// в то время как срезы работают с ним гораздо лучше. Следует
	// отметить, что это не проблема хеш-таблиц, а результат работы
	// сборщика мусора Go.
	// --- если не использовать хеш-таблицы, в которых хранятся огромные
	// объемы данных, в ваших программах эта проблема не будет столь
	// очевидной
}
