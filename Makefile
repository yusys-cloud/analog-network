.PHONY: all

all: test build run

.PHONY: run

run:
	@go run ./

.PHONY: build

build:
	@ rm -rf ./build/
	@go build -mod=vendor -o ./build/analog-network ./
	@echo "[OK] Server was build!"
	@cp config.json build
	@cp README.md  build
	@zip -r build/analog-network-mac.zip build

build-linux:
	@ rm -rf ./build/
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o ./build/analog-network-linux ./
	@echo "[OK] Linux Server was build!"
	@cp ./config.json ./build/
	@cp ./README.md  ./build/
	@zip -r ./build/analog-network-linux.zip ./build

build-windows:
	@ rm -rf ./build/
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -mod=vendor -o ./build/analog-network-linux ./
	@echo "[OK] Linux Server was build!"
	@cp ./config.json ./build/
	@cp ./README.md  ./build/

.PHONY: test

test:
	@go test -v -coverprofile=cover.out ./
	@echo "[OK] Test and coverage file was created!"

.PHONY: show_coverage

show_coverage:
	@go tool cover -html=cover.out
	@echo "[OK] Coverage file opened!"