package main

import (
	"fmt"
	"net/http"
	"time"
)

func aHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
	PORT := ":8001"

	// все URL-адреса, кроме /time, обслуживаются функцией aHandler(),
	// поскольку ее первый аргумент, равный /, соответствует любому
	// URL-адресу, который не подходит никакому другому обработчику
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/", aHandler)

	// Чтобы запустить веб-сервер, необходимо воспользоваться функцией
	// http.ListenAndServe(), указав соответствующий номер порта
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
