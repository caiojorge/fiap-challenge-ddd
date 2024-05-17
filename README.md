# Project Name

Projeto para finalização do curso de DDD, das disciplinas de DDD, Docker e Arquitetura 

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Design](#Design)
- [Deliverables](#Deliverablees)

## Installation

# SWAG no desenvolvimento apenas.
- Se o comando swag --version não funcionar, executar os passos abaixo:
    Para gerar a documentação no padrão open api, será necessário instalar o swag

go install github.com/swaggo/swag/cmd/swag@latest
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
swag --version


## Usage

- Instructions on how to use the project and any relevant examples.

- Como o CPF precisa ser válido, pf, busque um cpf no site: https://geradordecpf.dev/

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