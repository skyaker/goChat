// package main

// import (
// 	"fmt"
// 	routes "messages_service/src/routes"
// 	"net/http"

// 	_ "github.com/lib/pq"
// )

// func main() {
// 	http.HandleFunc("/send", routes.CreateMessage)
// 	fmt.Println("Server was started at http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	wrappedMux := loggingMiddleware(mux)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", wrappedMux)
}
