MGOOS=windows GOARCH=amd64AIN_PACKAGE_PATH := ./*.go
BINARY_NAME := argo
WINDOWS_BINARY_NAME := argo.exe

# ==============================================================================================#
# HELPERS
# ==============================================================================================#

## help: print the help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

no-dirty:
	git diff --exit-code

# ==============================================================================================#
# QUALITY CONTROL
# ==============================================================================================#


## tidy: format code and tidy modfile
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...

# ==============================================================================================#
# DEVELOPMENT
# ==============================================================================================#


## test: run all tests
test:
	go test -v -race -buildvcs ./...


## test/cover: run all tests and display coverage
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## build: build the application
build:
	go build -o=./bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}
	GOOS=windows GOARCH=amd64 go build -o=./bin/${WINDOWS_BINARY_NAME} ${MAIN_PACKAGE_PATH}


## run: run the application
run: build
	./bin/${BINARY_NAME}

# ==============================================================================================#
# OPERATIONS
# ==============================================================================================#

## push: push changes to the remote Git repository
push: tidy audit no-dirty
	git push origin main
