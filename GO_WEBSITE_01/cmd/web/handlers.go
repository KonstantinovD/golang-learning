package main

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"leaning/GO_WEBSITE_01/pkg/models"
	"net/http"
	"strconv"
)

// Создается функция-обработчик "home", которая записывает байтовый
// слайс, содержащий текст

// Функция-обработчик home является обычной функцией с двумя параметрами
// Параметр http.ResponseWriter предоставляет методы для объединения
// HTTP ответа и возвращение его пользователю, а второй параметр
// *http.Request является указателем на структуру, которая содержит
// информацию о текущем запросе (POST, GET, DELETE…, URL запроса)
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Проверяется, если текущий путь URL запроса точно совпадает с
	// шаблоном "/". Если нет, вызывается функция http.NotFound()
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippet := range s {
		fmt.Fprintf(w, "%v\n", snippet)
	}
	//
	//files := []string{
	//	"./GO_WEBSITE_01/ui/html/home.page.gohtml",
	//	"./GO_WEBSITE_01/ui/html/base.layout.gohtml",
	//	"./GO_WEBSITE_01/ui/html/footer.partial.gohtml",
	//}
	//// Используем функцию template.ParseFiles() для чтения файла шаблона.
	//// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	//// используя функцию http.Error() мы отправим пользователю
	//// ответ: 500 Internal Server Error
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.serverError(w, err)
	//	return
	//}
	//// Затем мы используем метод Execute() для записи содержимого
	//// шаблона в тело HTTP ответа. Последний параметр в Execute()
	//// предоставляет возможность отправки динамических данных в шаблон.
	//err = ts.Execute(w, nil)
	//if err != nil {
	//	app.serverError(w, err)
	//	return
	//}
}

// Обработчик для отображения содержимого заметки.
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// Извлекаем значение параметра id из URL и попытаемся
	// конвертировать строку в integer. Если его нельзя конвертировать,
	// или значение меньше 1, возвращаем 404
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// Используем функцию fmt.Fprintf() для вставки значения из id в строку ответа
	// и записываем его в http.ResponseWriter.
	fmt.Fprintf(w, "%v", s)
}

// Обработчик для создания новой заметки.
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
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
func (app *application) createEmptySnippet(w http.ResponseWriter, r *http.Request) {
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

func (app *application) showBiggestId(
	w http.ResponseWriter, r *http.Request) {
	app.snippets.Insert(models.Snippet{})
	w.Write([]byte("See app logs"))
}
