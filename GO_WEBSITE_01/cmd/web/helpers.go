package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// serverError записывает сообщение об ошибке в errorLog и
// затем отправляет пользователю error 500
func (app *application) serverError(w http.ResponseWriter, err error) {
	// используем функцию debug.Stack(), чтобы получить трассировку
	// стека для текущей горутины и добавить ее в логгер.
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	// Используем функцию http.StatusText() для текстового
	// представления кода состояния HTTP. К примеру,
	//http.StatusText(400) вернет строку "Bad Request"
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
