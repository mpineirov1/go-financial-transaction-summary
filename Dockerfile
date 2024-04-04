# Usa la imagen oficial de Golang como base
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /go/src/app

# Copia el código fuente de la aplicación al contenedor
COPY . .

# Compila la aplicación
RUN go build -o main .

# Comando por defecto para ejecutar la aplicación al iniciar el contenedor
CMD ["./main"]
