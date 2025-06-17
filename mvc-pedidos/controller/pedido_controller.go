package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"mvc-pedidos/model"
	"mvc-pedidos/view"

	"github.com/gorilla/mux"
)

// GET /pedidos
func ListarPedidos(w http.ResponseWriter, r *http.Request) {
	pedidos, err := model.BuscarTodosPedidos()
	if err != nil {
		view.RespondError(w, http.StatusInternalServerError, "Erro ao buscar pedidos")
		return
	}
	view.RespondJSON(w, http.StatusOK, pedidos)
}

// GET /pedidos/{id}
func BuscarPedidoPorID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	pedido, err := model.BuscarPedidoPorID(id)
	if err != nil {
		view.RespondError(w, http.StatusInternalServerError, "Erro ao buscar pedido")
		return
	}
	if pedido == nil {
		view.RespondError(w, http.StatusNotFound, "Pedido não encontrado")
		return
	}
	view.RespondJSON(w, http.StatusOK, pedido)
}

// GET /pedidos/carrinho/{id_carrinho}
func BuscarPorCarrinho(w http.ResponseWriter, r *http.Request) {
	idCarrinho := mux.Vars(r)["id_carrinho"]

	pedidos, err := model.BuscarPedidosPorCarrinho(idCarrinho)
	if err != nil {
		view.RespondError(w, http.StatusInternalServerError, "Erro ao buscar pedidos por carrinho")
		return
	}
	view.RespondJSON(w, http.StatusOK, pedidos)
}

// POST /pedidos
func CriarPedido(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var pedido model.Pedido
	err := json.Unmarshal(body, &pedido)
	if err != nil {
		view.RespondError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	id, err := model.CriarPedido(pedido)
	if err != nil {
		view.RespondError(w, http.StatusInternalServerError, "Erro ao criar pedido")
		return
	}

	view.RespondJSON(w, http.StatusCreated, map[string]any{"id": id})
}

// PUT /pedidos/{id}
func AtualizarPedido(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	body, _ := ioutil.ReadAll(r.Body)
	var pedido model.Pedido
	err := json.Unmarshal(body, &pedido)
	if err != nil {
		view.RespondError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	err = model.AtualizarPedido(id, pedido)
	if err != nil {
		view.RespondError(w, http.StatusInternalServerError, "Erro ao atualizar pedido")
		return
	}

	view.RespondJSON(w, http.StatusOK, map[string]string{"mensagem": "Pedido atualizado com sucesso"})
}

// DELETE /pedidos/{id}
func DeletarPedido(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)

	err := model.DeletarPedido(id)
	if err != nil {
		view.RespondError(w, http.StatusInternalServerError, "Erro ao deletar pedido")
		return
	}

	view.RespondJSON(w, http.StatusOK, map[string]string{"mensagem": fmt.Sprintf("Pedido %d excluído", id)})
}
