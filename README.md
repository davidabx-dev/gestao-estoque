# 📦 Gestão de Estoque (Golang + Clean Architecture)

Este projeto é uma implementação de referência de uma **API RESTful** em **Go**, focada em boas práticas de engenharia de software, desacoplamento e testabilidade.

O objetivo foi atender aos requisitos de uma vaga sênior, demonstrando domínio sobre **Clean Architecture**, **DDD (Domain-Driven Design)** e **Containerização**.

## 🏗️ Destaques da Arquitetura

O projeto não é apenas um CRUD simples. Ele foi estruturado para garantir que as regras de negócio independam de frameworks ou bancos de dados:

* **Clean Architecture:** Separação rígida entre camadas (Entities, UseCases, Controllers, Infra).
* **DDD (Domain-Driven Design):** Foco no domínio (Core) da aplicação.
* **Repository Pattern:** Desacoplamento do banco de dados usando interfaces.
* **Dependency Injection:** Injeção manual de dependências para facilitar testes.
* **TDD (Test Driven Development):** Testes unitários cobrindo as regras de negócio.

## 🚀 Stack Tecnológica

* **Linguagem:** Go (Golang) 1.25+
* **Banco de Dados:** SQL Nativo (Driver `database/sql` genérico, fácil migração para Postgres/MySQL).
* **Infraestrutura:** Docker (Multistage Build para imagens Alpine leves).
* **Testes:** Go Testing Package.

## 📂 Estrutura de Pastas

```text
/interno
  /entity      # Regras de negócio puras (sem imports externos)
  /usecase     # Casos de uso da aplicação (orquestração)
  /infra       # Implementações externas (Banco de dados, HTTP)
/cmd           # Entrypoint (Main)
Dockerfile     # Configuração de container otimizada
```

---

## ⚙️ Como Executar

### Opção 1: Via Docker (Recomendado)
A aplicação utiliza **Multistage Build**, gerando um container final extremamente leve.

```bash
# 1. Construir a imagem
docker build -t gestao-estoque .

# 2. Rodar o container (Mapeando porta 8000)
docker run -p 8000:8000 gestao-estoque
```

---

**Opção 2: Rodar Localmente (Go instalado)**

```bash
go mod tidy
go run cmd/server/main.go
```

---

### 🧪 Testando a API

**Health Check**

```bash
curl http://localhost:8080/health
```

---

**Criar Produto (Exemplo)**

```bash
curl -X POST http://localhost:8000/products \
   -H "Content-Type: application/json" \
   -d '{"name": "Teclado Mecânico", "price": 150.00}'
```

---

### 👨‍💻 Autor

Desenvolvido com foco em **Performance** e **Arquitetura de Software**.
