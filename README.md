# Project Name

- Projeto para finalização do curso de DDD, das disciplinas de DDD, Docker e Arquitetura Hexagonal

- Estou usando a linguagem go para a resolução do exercício.

## Caso necessário, seguem os passos para instalação do SWAG

# SWAG no desenvolvimento apenas.
- Se o comando swag --version não funcionar, executar os passos abaixo:
    Para gerar a documentação no padrão open api, será necessário instalar o swag

go install github.com/swaggo/swag/cmd/swag@latest
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
swag --version


## Como testar os endpoints?

 # Server
 - Executar o comando: docker-compose up -d
 - Executar o comando: go run cmd/kitchencontrol/main.go 

 # Acessar o swagger
 http://localhost:8080/kitchencontrol/api/v1/docs/index.html

 # Products
 - Testar a rota de products:
 - - http://localhost:8080/kitchencontrol/api/v1/docs/index.html#/Products/post_products
 - A api espera o seguinte body: 
 
 {
  "category": "almoço",
  "description": "Massa com queijo",
  "name": "Massa",
  "price": 40
} 


## Design

- Estou seguindo o padrão indicado nas aulas e tbm estou seguindo as boas práticas de nomenclatura do go.
- Internal (o código fica protegido)
- CMD (o main file de projetos go ficam aqui por padrão)
- Estrutura de pastas seguindo o exemplo em TS que o professor disponibilizou

## Deliverables

- a. arquitetura hexagonal
- b. api 
    - 1. Cadastro do cliente (customer)
    - 2. Identificação do cliente via CPF (no produto?)   