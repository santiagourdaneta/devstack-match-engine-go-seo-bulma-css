package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := "8080"
	fs := http.FileServer(http.Dir("dist"))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Lógica para Pretty URLs: si no tiene extensión, intenta servir el .html
		path := r.URL.Path
		if path != "/" && !strings.Contains(path, ".") {
			if _, err := os.Stat("dist" + path + ".html"); err == nil {
				r.URL.Path += ".html"
			}
		}
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		fs.ServeHTTP(w, r)
	})

	log.Printf("📡 DevStack-Match activo en http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
