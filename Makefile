# Variables
BINARY_NAME=seo_generator
DIST_DIR=dist
DATA_FILE=data.json

.PHONY: all build run clean help

# Prepara el entorno desde cero
setup:
	@echo "📦 Inicializando módulos y dependencias..."
	go mod tidy
	@mkdir -p assets/fonts assets/img
	@echo "🚀 Todo listo. Ejecuta 'make all' para generar el sitio."

# Comando por defecto: Limpia, construye y genera todo
all: clean build run

## build: Compila el binario (ahora con gestión de módulos)
build:
	@echo "🔧 Compilando generador en Go..."
	go build -o $(BINARY_NAME) main.go

## run: Ejecuta el generador para crear HTMLs y Sitemap
run:
	@echo "🚀 Generando sitio estático y Sitemap..."
	./$(BINARY_NAME)
	@echo "✨ Proceso completado. Revisa la carpeta /$(DIST_DIR)"

## clean: Borra el binario y la carpeta de salida
clean:
	@echo "🧹 Limpiando archivos antiguos..."
	rm -f $(BINARY_NAME)
	rm -rf $(DIST_DIR)

## help: Muestra los comandos disponibles
help:
	@echo "Comandos disponibles:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' |  sed -e 's/^/ /'

## deploy: Ejecuta el script de despliegue
deploy:
	@bash deploy.sh

## lint: Ejecuta el formateador y el linter de Go
lint:
	@echo "🧼 Limpiando y analizando código..."
	gofmt -w .
	go vet ./...

## serve: Genera y levanta el servidor local usando Go
serve: all
	@go run cmd/server/main.go
