BINARY_NAME = Simpletop.exe
APP_NAME= Simpletop
VERSION = 0.1.0

build:
	DEL $(BINARY_NAME)
	fyne package -os windows -appID $(APP_NAME) -icon icon.png -name $(BINARY_NAME) -version $(VERSION)

run:
	go run .

clean:
	@echo "Cleaning..."
	@go clean
	@DEL $(BINARY_NAME)
	@echo "Clean complete."

test:
	@echo "Running tests..."
	@go test -v ./...
	@echo "Test complete."