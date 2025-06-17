# 📦 API de Pedidos (MVC - Golang + SQLite + Vue 3 SPA)

Projeto desenvolvido como entrega do desafio final do Bootcamp **Arquiteto(a) de Software**, focado em arquitetura MVC com implementação em **Go (net/http)**, banco de dados **SQLite3**, e interface cliente em **Vue 3 SPA**.

---

## 🎯 Objetivo
Criar uma API REST para gestão de **pedidos**, utilizando boas práticas arquiteturais (MVC, C4 Model), persistência local e entrega de uma interface SPA funcional para consumo dos endpoints.

---

## 🛠️ Tecnologias Utilizadas

| Camada        | Tecnologia                  |
|---------------|-----------------------------|
| Backend       | Go (net/http)               |
| Banco de Dados| SQLite3                     |
| ORM           | `database/sql` nativo       |
| Roteador      | Gorilla Mux                 |
| Frontend SPA  | Vue 3 + Axios (HTML standalone) |
| Diagramas     | Structurizr DSL (C4 Model)  |

---

## 🧩 Estrutura de Pastas
```
/mvc-pedidos/
├── main.go
├── database.sqlite
├── config/
│   └── db.go             # conexão SQLite
├── model/
│   └── pedido.go         # struct + métodos de acesso a dados
├── controller/
│   └── pedido_controller.go
├── view/
│   └── response.go       # manipuladores de resposta JSON
└── router/
    └── routes.go         # mapeamento das rotas

```

---

## 🔗 Endpoints da API REST

| Verbo | Rota                          | Ação                       |
|-------|-------------------------------|-------------------------------|
| GET   | `/pedidos`                    | Lista todos os pedidos        |
| GET   | `/pedidos/{id}`              | Busca um pedido por ID        |
| GET   | `/pedidos/carrinho/{id}`     | Lista pedidos por carrinho    |
| POST  | `/pedidos`                   | Cria um novo pedido           |
| PUT   | `/pedidos/{id}`              | Atualiza um pedido existente  |
| DELETE| `/pedidos/{id}`              | Exclui um pedido              |
| GET   | `/health`                    | Verifica se a API está online |

---

## 📊 Diagramas C4 (Structurizr DSL)

O projeto inclui todos os níveis do C4 Model:
- **Nível 1**: Contexto (Usuário -> API <- Sistemas Terceiros)
- **Nível 2**: Containers (Vue, API Go, SQLite)
- **Nível 3**: Componentes (MVC Interno)
- **Nível 4**: Fluxo dinâmico e arquitetura de classes

---

## 🚀 Executar o Projeto

### Backend (Go)
```bash
go mod tidy
go run .
```

### Frontend (SPA)
- Abrir `main.html` em um navegador
- Certifique-se de que a API está rodando em `http://localhost:8080`

---

## ⚠️ CORS - Importante para Desenvolvimento
Para evitar erros de CORS ao usar o SPA em `file://`, o backend inclui um middleware que permite qualquer origem (*).

---

## ✍️ Autor e Créditos
Desenvolvido por **Wilson Missina Faria** como parte do desafio final de arquitetura de software. Todos os diagramas foram gerados em [Structurizr DSL](https://structurizr.com/help/dsl) e o frontend em Html + Vue 3 + Axios.
