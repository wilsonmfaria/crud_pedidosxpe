package router

import (
	"net/http"

	"mvc-pedidos/controller"

	"github.com/gorilla/mux"
)

// Inicializa e retorna o roteador configurado
func InitRoutes() *mux.Router {
	rr := mux.NewRouter()

	// Cria um prefixo de rota para a API
	r := rr.PathPrefix("/api").Subrouter()

	// Endpoints de pedidos
	r.HandleFunc("/pedidos", controller.ListarPedidos).Methods("GET")
	r.HandleFunc("/pedidos/{id:[0-9]+}", controller.BuscarPedidoPorID).Methods("GET")
	r.HandleFunc("/pedidos/carrinho/{id_carrinho}", controller.BuscarPorCarrinho).Methods("GET")
	r.HandleFunc("/pedidos", controller.CriarPedido).Methods("POST")
	r.HandleFunc("/pedidos/{id:[0-9]+}", controller.AtualizarPedido).Methods("PUT")
	r.HandleFunc("/pedidos/{id:[0-9]+}", controller.DeletarPedido).Methods("DELETE")

	// Health check simples
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return r
}
