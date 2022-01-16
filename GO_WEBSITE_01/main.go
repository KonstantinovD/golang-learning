package main

import (
	"log"
	"net/http"
)

// Создается функция-обработчик "home", которая записывает байтовый
// слайс, содержащий текст

// Функция-обработчик home является обычной функцией с двумя параметрами
// Параметр http.ResponseWriter предоставляет методы для объединения
// HTTP ответа и возвращение его пользователю, а второй параметр
// *http.Request является указателем на структуру, которая содержит
// информацию о текущем запросе (POST, GET, DELETE…, URL запроса)
func home(w http.ResponseWriter, r *http.Request) {
	// Проверяется, если текущий путь URL запроса точно совпадает с
	// шаблоном "/". Если нет, вызывается функция http.NotFound()
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from golang-website"))
}

// Обработчик для отображения содержимого заметки.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Отображение заметки..."))
}

// Обработчик для создания новой заметки.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Форма для создания новой заметки..."))
}

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

	// Используется функция http.ListenAndServe() для запуска нового
	// веб-сервера. Мы передаем два параметра: TCP-адрес сети для
	// прослушивания и созданный рутер.
	log.Println("Запуск веб-сервера на http://127.0.0.1:8001")
	err := http.ListenAndServe(":8001", mux)
	log.Fatal(err)
}
