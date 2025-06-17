package model

import (
	"database/sql"
	"mvc-pedidos/config"
)

// Pedido representa um item de um carrinho de compras
type Pedido struct {
	ID         int    `json:"id"`
	IdCarrinho string `json:"id_carrinho"`
	IdUsuario  int    `json:"id_usuario"`
	IdProduto  int    `json:"id_produto"`
	Quantidade int    `json:"qdt"`
}

// Cria um novo pedido
func CriarPedido(p Pedido) (int64, error) {
	stmt, err := config.DB.Prepare(`INSERT INTO pedidos (id_carrinho, id_usuario, id_produto, qdt) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.IdCarrinho, p.IdUsuario, p.IdProduto, p.Quantidade)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// Lista todos os pedidos
func BuscarTodosPedidos() ([]Pedido, error) {
	rows, err := config.DB.Query(`SELECT id, id_carrinho, id_usuario, id_produto, qdt FROM pedidos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pedidos []Pedido
	for rows.Next() {
		var p Pedido
		if err := rows.Scan(&p.ID, &p.IdCarrinho, &p.IdUsuario, &p.IdProduto, &p.Quantidade); err != nil {
			return nil, err
		}
		pedidos = append(pedidos, p)
	}
	return pedidos, nil
}

// Busca pedido por ID
func BuscarPedidoPorID(id int) (*Pedido, error) {
	row := config.DB.QueryRow(`SELECT id, id_carrinho, id_usuario, id_produto, qdt FROM pedidos WHERE id = ?`, id)

	var p Pedido
	if err := row.Scan(&p.ID, &p.IdCarrinho, &p.IdUsuario, &p.IdProduto, &p.Quantidade); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

// Atualiza um pedido
func AtualizarPedido(id int, p Pedido) error {
	stmt, err := config.DB.Prepare(`UPDATE pedidos SET id_carrinho = ?, id_usuario = ?, id_produto = ?, qdt = ? WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.IdCarrinho, p.IdUsuario, p.IdProduto, p.Quantidade, id)
	return err
}

// Deleta um pedido
func DeletarPedido(id int) error {
	stmt, err := config.DB.Prepare(`DELETE FROM pedidos WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

// Busca pedidos por ID do carrinho
func BuscarPedidosPorCarrinho(idCarrinho string) ([]Pedido, error) {
	rows, err := config.DB.Query(`SELECT id, id_carrinho, id_usuario, id_produto, qdt FROM pedidos WHERE id_carrinho = ?`, idCarrinho)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pedidos []Pedido
	for rows.Next() {
		var p Pedido
		if err := rows.Scan(&p.ID, &p.IdCarrinho, &p.IdUsuario, &p.IdProduto, &p.Quantidade); err != nil {
			return nil, err
		}
		pedidos = append(pedidos, p)
	}
	return pedidos, nil
}
