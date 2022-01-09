package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

// Переменную, которая используется пакетом atomic, мы сделали
// глобальной, чтобы она была доступной из любой точки кода
var count int32

// Счетчик atomic, применяемый в этой программе, связан с глобальной
// переменной count и помогает подсчитать общее количество клиентов,
// которые обслужил веб-сервер.
func handleAll(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt32(&count, 1)
}

func getCounter(w http.ResponseWriter, r *http.Request) {
	temp := atomic.LoadInt32(&count)
	fmt.Println("Count:", temp)
	fmt.Fprintf(w, "<h1 align=\"center\">%d</h1>", count)
}

func main() {
	http.HandleFunc("/getCounter", getCounter)
	http.HandleFunc("/", handleAll)
	http.ListenAndServe(":8001", nil)
}
