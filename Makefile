MAIN_PKG = cmd/main.go

BUILD_DIR = dist

CONFIG_PATH = config.yaml

serve:
	go build -o dist/main.exe cmd/main.go
	dist/main.exe -config=config.yaml

build:
	docker-compose up --build epta-app
run:
	docker-compose up epta-app
swag:
	swag init -g internal/transport/rest/handler/handlers.go