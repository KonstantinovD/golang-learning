package main

import "net/http"

func (app *application) routes() *http.ServeMux {
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
	mux.HandleFunc("/", app.home)                        // non-fixed path (ends with '/')
	mux.HandleFunc("/snippet", app.showSnippet)          // fixed path
	mux.HandleFunc("/snippet/create", app.createSnippet) // fixed path
	mux.HandleFunc("/snippet/empty/create", app.createEmptySnippet)
	mux.HandleFunc("/insert-snippet", app.showBiggestId)

	// Инициализируем FileServer, он будет обрабатывать
	// HTTP-запросы к статическим файлам из папки "./ui/static".
	// Обратите внимание, что переданный в функцию http.Dir путь
	// является относительным корневой папке проект
	fileServer := http.FileServer(
		http.Dir("./GO_WEBSITE_01/ui/static/"))
	// Попытка настроить файловую систему для статик файлов и
	// возврата 404 привело к недоступности стилей для страницы.
	// Поэтому вернул обратно
	/*
		fileServer := http.FileServer(neuteredFileSystem{
		http.Dir("./static/")})
	*/

	// Используем функцию mux.Handle() для регистрации обработчика для
	// всех запросов, которые начинаются с "/static/". Мы убираем
	// префикс "/static" перед тем как запрос достигнет http.FileServer
	mux.Handle("/static/",
		http.StripPrefix("/static", fileServer))

	return mux
}
