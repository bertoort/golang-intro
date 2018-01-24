package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

type turtle struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Color string `json:"color"`
}

func main() {
	port := os.Getenv("PORT")
	router := chi.NewRouter()
	donatelo := turtle{Name: "Donatelo", Age: 100, Color: "purple"}
	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(donatelo)
	})
	http.ListenAndServe(":"+port, router)
}
