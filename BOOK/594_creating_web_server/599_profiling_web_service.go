package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"time"
)

// для того чтобы профилировать приложение Go с собственным
// HTTP-сервером, следует использовать стандартный пакет Go
// net/http/pprof
// для профилирования остальных видов приложений (not web)
// применяется стандартный Go-пакет runtime/pprof

// Заметьте, что для регистрации поддерживаемых программой путей
// используется переменная http.NewServeMux, поскольку для того, чтобы
// использовать http.NewServeMux, нужно задать конечные точки HTTP
// вручную. Мы можем определять подмножество поддерживаемых конечных
// точек HTTP. Если вы решите не использовать http.NewServeMux, то
// конечные точки HTTP будут зарегистрированы автоматически. Для этого
// придется импортировать пакет net/http/pprof, поставив перед ним
// символ '_'

func aHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler2(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
	PORT := ":8001"

	// not working at windows? (need some researches)
	r := http.NewServeMux()
	http.HandleFunc("/time", timeHandler2)
	http.HandleFunc("/", aHandler2)

	// -- В этом коде Go определены конечные точки HTTP, связанные с
	// профилированием. Без них мы не сможем профилировать
	// веб-приложение
	// --- Здесь не определена конечная точка HTTP      |
	// /debug/pprof/goroutine                           V
	// Это допустимо, поскольку в данной проге отсутствуют горутины
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// второй параметр функции http.ListenAndServe() уже не равен nil
	err := http.ListenAndServe(PORT, r)
	if err != nil {
		fmt.Println(err)
		return
	}
}
