package main

import (
	"fmt"
	"log"
	"net/http"

	"mvc-pedidos/config"
	"mvc-pedidos/router"
)

func enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Permite qualquer origem (ideal para dev)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Tratamento especial para pré-flight OPTIONS
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}


func main() {
	fmt.Println("Iniciando aplicação MVC Pedidos em Go...")

	// Inicializa banco de dados e seed
	config.InitDB()
	defer config.DB.Close()

	// Inicializa rotas
	r := router.InitRoutes()

	// Sobe servidor HTTP
	porta := ":8080"
	fmt.Printf("Servidor escutando em http://localhost%s\n", porta)
	log.Fatal(http.ListenAndServe(porta, enableCORS(r)))
}
