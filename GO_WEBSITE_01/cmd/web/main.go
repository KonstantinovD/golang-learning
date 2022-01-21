package main

import (
	"log"
	"net/http"
	"path/filepath"
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

	// Инициализируем FileServer, он будет обрабатывать
	// HTTP-запросы к статическим файлам из папки "./ui/static".
	// Обратите внимание, что переданный в функцию http.Dir путь
	// является относительным корневой папке проект
	fileServer := http.FileServer(neuteredFileSystem{
		http.Dir("./static/")})
	// Используем функцию mux.Handle() для регистрации обработчика для
	// всех запросов, которые начинаются с "/static/". Мы убираем
	// префикс "/static" перед тем как запрос достигнет http.FileServer
	mux.Handle("/static/",
		http.StripPrefix("/static", fileServer))

	// Используется функция http.ListenAndServe() для запуска нового
	// веб-сервера. Мы передаем два параметра: TCP-адрес сети для
	// прослушивания и созданный рутер.
	log.Println("Запуск веб-сервера на http://127.0.0.1:8001")
	err := http.ListenAndServe(":8001", mux)
	log.Fatal(err)
}

// Более сложным способом (однако он считается лучшим) является создание
// настраиваемой имплементации файловой системы http.FileSystem, с
// помощью которой будет возвращаться ошибка os.ErrNotExist для любого
// HTTP запроса напрямую к папке.
//
// Например, если пользователь попытается открыть в браузере ссылку
// http://127.0.0.1:8001/static/ в браузере, то он не должен увидеть
// список файлов. Он должен получить ошибку '404'
type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}
	return f, nil
}
