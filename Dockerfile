# Usar una imagen oficial de Go como imagen base
FROM golang:1.21

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos Go y los módulos de Go al contenedor
COPY . .

# Descargar las dependencias (si utilizas Go Modules, que es el enfoque recomendado)
RUN go mod download

# Compilar la aplicación
RUN go build -o math-services main.go

# Exponer el puerto por el que el microservicio será accesible
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./math-services"]
