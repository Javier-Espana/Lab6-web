# Usar una imagen base de Debian para construir Go 1.23.0
FROM debian:bullseye

# Instalar dependencias necesarias
RUN apt-get update && apt-get install -y wget git build-essential

# Descargar y construir Go 1.23.0
RUN wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz && \
    rm go1.23.0.linux-amd64.tar.gz

# Configurar PATH para Go y herramientas instaladas
ENV PATH="/usr/local/go/bin:/root/go/bin:${PATH}"

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos del backend al contenedor
COPY . .

# Instalar las herramientas necesarias (como swaggo para Swagger)
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Descargar las dependencias del proyecto
RUN go mod tidy

# Generar la documentación de Swagger
RUN swag init

# Construir el ejecutable
RUN go build -o main .

# Listar los archivos en el directorio de trabajo
RUN ls -l /app

# Comando para ejecutar la aplicación
CMD ["./main"]