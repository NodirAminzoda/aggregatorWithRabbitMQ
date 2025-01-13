# Указываем базовый образ
FROM golang:1.22-alpine AS builder

ENV GOOS=linux

# Устанавливаем рабочую директорию
WORKDIR /app


# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код в рабочую директорию
COPY . .

# Собираем приложение
RUN go build -o myapp ./cmd/main.go

# Создаем финальный образ
FROM alpine

# Устанавливаем рабочую директорию
WORKDIR /app

RUN ls -la

# Копируем бинарник из предыдущего образа
COPY --from=builder /app/myapp .

# Копируем папку с конфигурациями
COPY --from=builder /app/internal/config/config.yaml /app/internal/config/config.yaml

EXPOSE 8001
# Указываем команду, которая будет выполнена при запуске контейнера
CMD ["./myapp"]
