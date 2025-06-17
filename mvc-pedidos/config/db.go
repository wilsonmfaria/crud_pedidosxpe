package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error

	// Cria o banco se não existir
	if _, err := os.Stat("database.sqlite"); os.IsNotExist(err) {
		file, err := os.Create("database.sqlite")
		if err != nil {
			log.Fatal("Erro ao criar banco de dados:", err)
		}
		file.Close()
		fmt.Println("Banco de dados criado com sucesso.")
	}

	DB, err = sql.Open("sqlite", "database.sqlite")
	if err != nil {
		log.Fatal("Erro ao abrir conexão com banco:", err)
	}

	createTables()
	seedData()
}

func createTables() {
	createUsuario := `
	CREATE TABLE IF NOT EXISTS usuarios (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nome TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);`

	createProduto := `
	CREATE TABLE IF NOT EXISTS produtos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nome TEXT NOT NULL,
		preco REAL NOT NULL,
		quantidade INTEGER NOT NULL
	);`

	createPedido := `
	CREATE TABLE IF NOT EXISTS pedidos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		id_carrinho TEXT NOT NULL,
		id_usuario INTEGER,
		id_produto INTEGER,
		qdt INTEGER NOT NULL,
		FOREIGN KEY(id_usuario) REFERENCES usuarios(id),
		FOREIGN KEY(id_produto) REFERENCES produtos(id)
	);`

	_, err := DB.Exec(createUsuario)
	checkError("usuarios", err)

	_, err = DB.Exec(createProduto)
	checkError("produtos", err)

	_, err = DB.Exec(createPedido)
	checkError("pedidos", err)
}

func seedData() {
	// verifica se já tem usuários
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM usuarios").Scan(&count)
	if err != nil || count > 0 {
		return
	}

	fmt.Println("Inserindo dados de exemplo...")

	tx, err := DB.Begin()
	if err != nil {
		log.Fatal("Erro ao iniciar transação de seed:", err)
	}

	_, err = tx.Exec(`INSERT INTO usuarios (nome, email) VALUES 
		('João Silva', 'joao@exemplo.com'),
		('Maria Oliveira', 'maria@exemplo.com');`)

	_, err = tx.Exec(`INSERT INTO produtos (nome, preco, quantidade) VALUES 
		('Caneta Azul', 1.99, 100),
		('Caderno 200 folhas', 15.90, 50),
		('Lápis Preto', 0.89, 200);`)

	if err != nil {
		tx.Rollback()
		log.Fatal("Erro ao inserir dados de exemplo:", err)
	}

	tx.Commit()
}

func checkError(entity string, err error) {
	if err != nil {
		log.Fatalf("Erro ao criar tabela %s: %v", entity, err)
	}
}
