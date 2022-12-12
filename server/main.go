package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println("Iniciou minha request")

	defer log.Println("Finalizou minha request")

	select {
	case <-time.After(time.Second * 5):
		log.Println("Requisição realizada com sucesso")
		w.Write([]byte("Requisição realizada com sucesso - http"))
	case <-ctx.Done():
		log.Println("Request cancelada")
		http.Error(w, "Request cancelada - http", http.StatusRequestTimeout)
	}
}
