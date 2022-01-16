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
	// -- Используем метод Header().Set() для добавления заголовка
	// 'Allow: POST' в map HTTP-headers, чтобы пользователь знал,
	// какие HTTP-методы поддерживаются для определенного URL.
	// -- Header().Set вызывается строго перед WriteHeader()/Write()
	w.Header().Set("Allow", http.MethodPost)
	// Используем r.Method для проверки, использует ли запрос метод POST
	// или нет. Обратите внимание, что http.MethodPost является строкой
	// и содержит текст "POST".
	if r.Method != http.MethodPost {
		// Если это не так, то вызывается метод w.WriteHeader() для
		// возврата статус-кода 405 и вызывается метод w.Write() для
		// тела-ответа с текстом "Метод запрещен".
		w.WriteHeader(405)
		// -- Вызвать метод w.WriteHeader() в обработчике можно только
		// один раз, иначе Go выдаст сообщение об ошибке;
		// -- Если не вызывать метод w.WriteHeader() напрямую, тогда
		// первый вызов w.Write() автоматически отправит пользователю
		// код состояния 200 OK. Поэтому, если вы хотите вернуть другой
		// код состояния, вызовите один раз метод w.WriteHeader() перед
		// любым вызовом w.Write().
		w.Write([]byte("GET-Метод запрещен!"))
		return
	}

	w.Write([]byte("Форма для создания новой заметки..."))
}

// Обработчик для создания пустой заметки.
func createEmptySnippet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", http.MethodPost)
	if r.Method != http.MethodPost {
		// Если требуется отправить какой-то код состояния, кроме 200
		// с текстом ответа, можно использовать http.Error(). Эта
		// функция принимает текст сообщения и код состояния, а затем
		// сама вызывает методы w.WriteHeader() и w.Write()
		http.Error(w, "Метод запрещен!", 405)
		return
	}

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
	mux.HandleFunc("/snippet/empty/create", createEmptySnippet)

	// Используется функция http.ListenAndServe() для запуска нового
	// веб-сервера. Мы передаем два параметра: TCP-адрес сети для
	// прослушивания и созданный рутер.
	log.Println("Запуск веб-сервера на http://127.0.0.1:8001")
	err := http.ListenAndServe(":8001", mux)
	log.Fatal(err)
}
