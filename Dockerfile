# Usa una imagen base de Go
FROM golang:1.20-alpine

# Define el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el código fuente al contenedor
COPY . .

# Compila el proyecto
RUN go mod init WebHook.go
RUN go mod tidy
RUN go build -o WebHook.go .

# Expone el puerto 8080 para recibir peticiones
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./WebHook-go"]
