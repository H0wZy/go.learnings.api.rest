# 🚀 go.learnings.api.rest

> API REST em Go para gerenciamento de vagas de emprego, desenvolvida com foco em aprendizado e boas práticas.

[![Go Version](https://img.shields.io/badge/Go-1.25.3-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/Gin-v1.11.0-00ADD8?style=flat)](https://gin-gonic.com/)
[![SQLite](https://img.shields.io/badge/SQLite-3-003B57?style=flat&logo=sqlite)](https://www.sqlite.org/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)](https://www.docker.com/)

## 📚 Sobre o Projeto

Este projeto foi desenvolvido com o intuito de **aprender e estudar API REST** e **boas práticas de estruturação e convenção de projetos em Go**, com base nos ensinamentos de [arthur404dev](https://github.com/arthur404dev).
Assista o [vídeo completo](https://www.youtube.com/watch?v=wyEYpX5U4Vg) dele no YouTube.

A aplicação implementa um CRUD completo para gerenciamento de vagas de emprego (job openings), utilizando tecnologias modernas e seguindo padrões da comunidade Go.

## ✨ Funcionalidades

- ✅ **CRUD Completo** de vagas de emprego
- 📝 **Documentação automática** com Swagger/OpenAPI
- 🗄️ **Banco de dados SQLite** com GORM
- 🔍 **Logging estruturado** customizado
- 🐳 **Docker & Docker Compose** para ambientes isolados
- ⚡ **Hot reload** no ambiente de desenvolvimento
- 🎯 **API RESTful** seguindo boas práticas

## 🏗️ Estrutura do Projeto

```
.
├── config/          # Configurações da aplicação
│   ├── config.go    # Configuração geral
│   ├── logger.go    # Logger customizado
│   └── sqlite.go    # Configuração do banco de dados
├── db/              # Arquivos de banco de dados
├── docs/            # Documentação Swagger gerada
├── handler/         # Controllers/Handlers da API
│   ├── createOpening.go
│   ├── deleteOpening.go
│   ├── listOpenings.go
│   ├── showOpening.go
│   ├── updateOpening.go
│   ├── request.go   # Validação de requests
│   └── response.go  # Formatação de responses
├── router/          # Configuração de rotas
│   ├── router.go
│   └── routes.go
├── schemas/         # Modelos de dados
│   └── opening.go
├── .env             # Variáveis de ambiente
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── main.go
├── makefile         # Comandos úteis
└── README.md
```

## 🛠️ Tecnologias

- **[Go 1.25.3](https://golang.org/)** - Linguagem de programação
- **[Gin](https://gin-gonic.com/)** - Framework web
- **[GORM](https://gorm.io/)** - ORM para Go
- **[SQLite](https://www.sqlite.org/)** - Banco de dados
- **[Swag](https://github.com/swaggo/swag)** - Geração de documentação Swagger
- **[Docker](https://www.docker.com/)** - Containerização

## 🚀 Como Executar

### Pré-requisitos

- Go 1.25.3+
- Docker & Docker Compose (opcional)
- Make (opcional, mas recomendado)

### Configuração

1. **Clone o repositório:**
```bash
git clone https://github.com/H0wZy/go.learnings.api.rest.git
cd go.learnings.api.rest
```

2. **Configure as variáveis de ambiente:**
```bash
cp .env.example .env
# Edite o .env conforme necessário
```

3. **Instale as dependências:**
```bash
go mod download
```

### Executando Localmente

**Com Make (recomendado):**
```bash
# Executar com documentação Swagger
make runwdocs

# Apenas executar
make run

# Build da aplicação
make gobuild

# Executar testes
make test

# Gerar documentação Swagger
make docs

# Limpar arquivos gerados
make cls
```

**Sem Make:**
```bash
# Gerar docs e executar
swag init && go run main.go

# Apenas executar
go run main.go
```

### Executando com Docker

**Com Make:**
```bash
# Build da imagem
make build

# Subir o container
make up

# Ver logs
make logs

# Parar o container
make stop

# Remover o container
make down
```

**Sem Make:**
```bash
docker compose build
docker compose up -d
docker compose logs -f
docker compose down
```

## 📡 Endpoints da API

### Base URL: `http://localhost:8080/api/v1`

| Método   | Endpoint        | Descrição |
|----------|-----------------|-----------|
| `POST`   | `/opening`      | Criar uma nova vaga |
| `GET`    | `/openings`     | Listar todas as vagas |
| `GET`    | `/opening?id=0` | Buscar vaga por ID |
| `PATCH`  | `/opening?id=0` | Atualizar uma vaga |
| `DELETE` | `/opening?id=0` | Deletar uma vaga |

### Exemplo de Request (POST)

```json
{
  "role": "Backend Developer",
  "company": "Tech Corp",
  "location": "São Paulo, SP",
  "remote": true,
  "link": "https://techcorp.com/jobs/backend",
  "salary": 8000
}
```

### Exemplo de Response

```json
{
  "data": {
    "id": 1,
    "createdAt": "2024-11-09T20:00:00Z",
    "updatedAt": "2024-11-09T20:00:00Z",
    "role": "Backend Developer",
    "company": "Tech Corp",
    "location": "São Paulo, SP",
    "remote": true,
    "link": "https://techcorp.com/jobs/backend",
    "salary": 8000
  },
  "message": "opening created successfully"
}
```

## 📖 Documentação Swagger

Após iniciar a aplicação, acesse:

```
http://localhost:8080/swagger/index.html
```

## ⚙️ Variáveis de Ambiente

| Variável | Descrição | Padrão |
|----------|-----------|--------|
| `APP_NAME` | Nome da aplicação | `go.learnings.api.rest` |
| `PORT` | Porta da aplicação | `8080` |
| `GIN_MODE` | Modo do Gin (debug/release) | `release` |
| `GO_VERSION` | Versão do Go (Docker) | `1.25.3-alpine` |
| `DB_PATH` | Caminho do banco SQLite | `./db/main.db` |
| `CONTAINER_NAME` | Nome do container Docker | `go-api-rest` |
| `IMAGE_NAME` | Nome da imagem Docker | `go.learnings.api.rest` |
| `IMAGE_TAG` | Tag da imagem Docker | `latest` |

## 🧪 Testes

```bash
make test
# ou
go test ./...
```

## 📝 Comandos Make

| Comando | Descrição |
|---------|-----------|
| `make` | Executa `runwdocs` (padrão) |
| `make run` | Executa a aplicação |
| `make runwdocs` | Gera docs e executa |
| `make gobuild` | Compila a aplicação |
| `make test` | Executa os testes |
| `make docs` | Gera documentação Swagger |
| `make cls` | Remove arquivos gerados |
| `make build` | Build da imagem Docker |
| `make up` | Sobe o container |
| `make stop` | Para o container |
| `make down` | Remove o container |
| `make logs` | Mostra logs do container |

## 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.

1. Fork o projeto
2. Crie sua feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença Apache 2.0. Veja o arquivo [LICENSE](LICENSE.md) para mais detalhes.

## 👤 Autor

**H0wZy**

- GitHub: [@H0wZy](https://github.com/H0wZy)

## 🙏 Agradecimentos

- [arthur404dev](https://github.com/arthur404dev) - Pelos ensinamentos e boas práticas
- Comunidade Go Brasil
- Documentação oficial do Go e Gin Framework

---

⭐ Se este projeto te ajudou, considere dar uma estrela!
