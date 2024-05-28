# Usar a imagem oficial do Go para construir a aplicação
FROM golang:1.22.1 AS builder

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Copiar go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod download

# Copiar o código da aplicação para o diretório de trabalho
COPY . .

# Compilar a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fiap-rocks ./cmd/kitchencontrol/main.go

# Usar uma imagem mais leve para rodar a aplicação
FROM alpine:latest

# Instalar as dependências necessárias
RUN apk --no-cache add ca-certificates

# Definir o diretório de trabalho dentro do container
WORKDIR /root/

# Copiar o binário compilado da imagem builder
COPY --from=builder /app/fiap-rocks .

# Certificar que o binário tem permissões de execução
RUN chmod +x ./fiap-rocks

# Expor a porta em que a aplicação vai rodar
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./fiap-rocks"]
