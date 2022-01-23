package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Создаем структуру `application` для хранения зависимостей всего
// веб-приложения. Пока, что мы добавим поля только для двух логгеров,
// мы будем расширять данную структуру по мере усложнения приложения.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	// Используйте log.New() для создания логгера. Для этого нужно
	// три параметра: место назначения для записи логов (os.Stdout),
	// строка с префиксом сообщения (INFO или ERROR) и флаги,
	// указывающие, какая дополнительная информация будет добавлена.
	// Обратите внимание - флаги соединяются с помощью оператора OR |
	infoLog := log.New(os.Stdout, "INFO\t",
		log.Ldate|log.Ltime)
	//
	// используем флаг log.Lshortfile для включения в лог
	// названия файла и номера строки где обнаружилась ошибка
	errorLog := log.New(os.Stderr, "ERROR\t",
		log.Ldate|log.Ltime|log.Lshortfile)

	// Структура с зависимостями приложения.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	// рутер
	mux := app.routes()

	// Используется функция http.ListenAndServe() для запуска нового
	// веб-сервера. Мы передаем два параметра: TCP-адрес сети для
	// прослушивания и созданный рутер.
	infoLog.Println("Запуск веб-сервера на http://127.0.0.1:8001")
	err := http.ListenAndServe(":8001", mux)
	errorLog.Fatal(err)
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
