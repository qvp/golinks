# Используем официальный образ Go в качестве базового образа
FROM golang:1.24-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum (если есть) для загрузки зависимостей
COPY go.mod ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код в контейнер
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o golinks .

# Используем минимальный образ alpine для финального контейнера
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем собранный бинарный файл из builder stage
COPY --from=builder /app/golinks .

# Открываем порт, на котором будет работать приложение
#EXPOSE 8080

# Запускаем приложение
#CMD ["./golinks"]