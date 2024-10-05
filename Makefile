MAIN_PKG = cmd/main.go

BUILD_DIR = dist

CONFIG_PATH = config.yaml

serve:
	go build -o ${BUILD_DIR}/main.exe ${MAIN_PKG}
	${BUILD_DIR}/main.exe -config=${CONFIG_PATH}

build:
	docker-compose up --build epta-app
run:
	docker-compose up epta-app
swag:
	swag init -g internal/transport/rest/handler/handlers.go