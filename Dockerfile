# --- ETAPA 1: Builder (Compilação) ---
FROM golang:1.25 as builder

WORKDIR /app

# Copia os arquivos de dependência e baixa
COPY go.mod go.sum ./
RUN go mod download

# Copia o resto do código
COPY . .

# Compila um binário estático (sem dependências de sistema) para Linux
# CGO_ENABLED=0 é vital para rodar numa imagem vazia (scratch)
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/server/main.go

# --- ETAPA 2: Runner (Imagem Final) ---
FROM scratch

WORKDIR /app

# Copia apenas o executável da etapa anterior
COPY --from=builder /app/server .

# Expõe a porta
EXPOSE 8000

# Comando para rodar
CMD ["./server"]