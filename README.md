# Projeto para finalização do curso de DDD, das disciplinas de DDD, Docker e Arquitetura Hexagonal

## Descrição

A Kitchen Control API é uma aplicação para gerenciar clientes, produtos, pedidos e itens de pedidos. Esta API fornece endpoints para criar, buscar, atualizar e deletar registros.

## Tecnologias

- Go
- GORM
- MySQL
- Gin Web Framework

## Instalação

### Pré-requisitos

- wsl 2, macos ou linux
- Go 1.22.1 ou superior
- Docker
- Git
- make

### Passos para Instalação

1. Clone o repositório:
    ```bash
    git clone https://github.com/caiojorge/fiap-challenge-ddd.git
    cd fiap-challenge-ddd
    ```

2. Instale as dependências:
    ```bash
    go mod tidy
    ```

3. Execute as migrações do banco de dados:
    ```bash
    make mysql (ou docker-compose up -d)
    make run (ou go run cmd/kitchencontrol/main.go)
    ```
- O arquivo init-db esta conectado ao docker, e deve ser executado para criar o banco de dados caso não exista

4. Execute os testes:
    ```bash
    make test (ou go test -v -cover ./...)
    ```

4. Inicie o servidor e acesse o swagger:
    ```bash
    make run (ou go run cmd/kitchencontrol/main.go)
    http://localhost:8080/kitchencontrol/api/v1/docs/index.html
    ```

## Uso

### Endpoints (acesso via swagger)
- http://localhost:8080/kitchencontrol/api/v1/docs/index.html
- Acessar o swagger, e toda documentação de uso está lá.


## Documentação
- O link para acesso ao miro será enviado aos professores via plataforma da fiap

## Licença
Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## SWAG no desenvolvimento apenas.
- Se o comando swag --version não funcionar, executar os passos abaixo:
    - Para gerar a documentação no padrão open api, será necessário instalar o swag

    ```bash
    go install github.com/swaggo/swag/cmd/swag@latest
    echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
    source ~/.bashrc
    swag --version
    ```
    - no macos, ao invés de .bashrc, use .zhrc
