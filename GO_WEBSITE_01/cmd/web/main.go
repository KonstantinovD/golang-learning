package main

import (
	"log"
	"net/http"
)

func main() {
	// -- Используется функция http.NewServeMux() для инициализации
	// нового рутера, затем функции регистрируются как обработчики
	// -- При работе с Go встречались функции http.Handle()
	// и http.HandleFunc(). Они регают пути с помощью DefaultServeMux.
	// Мы используем собственный ServeMux, поскольку DefaultServeMux
	// является глобальной переменной и любой пакет может получить к ней
	// доступ и зарегистрировать маршрут — включая любые сторонние
	// пакеты, которые использует ваше приложение.
	mux := http.NewServeMux()
	// шаблон "/" действует по сценарию «catch-all»
	mux.HandleFunc("/", home)                        // non-fixed path (ends with '/')
	mux.HandleFunc("/snippet", showSnippet)          // fixed path
	mux.HandleFunc("/snippet/create", createSnippet) // fixed path
	mux.HandleFunc("/snippet/empty/create", createEmptySnippet)

	// Используется функция http.ListenAndServe() для запуска нового
	// веб-сервера. Мы передаем два параметра: TCP-адрес сети для
	// прослушивания и созданный рутер.
	log.Println("Запуск веб-сервера на http://127.0.0.1:8001")
	err := http.ListenAndServe(":8001", mux)
	log.Fatal(err)
}
