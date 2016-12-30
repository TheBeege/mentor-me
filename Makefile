SSHKEY=""
DEPLOY_HOST="localhost"

REPO="github.com/TheBeege/mentor-me"
BINARY_NAME="mentor-me"
BRANCH=`git rev-parse --abbrev-ref HEAD`
COMMIT=`git rev-parse --short HEAD`
GOLDFLAGS="-X main.branch=$(BRANCH) -X main.commit=$(COMMIT)"
DBUSER=postgres

PACKAGES="models"

all: test build dump_schema

setup:
	@echo "=== setup ==="
	@echo "Make sure to install Mercurial!"
	@go get -u "github.com/tools/godep"
	@go get -u "github.com/golang/lint/golint"
	@go get -u "github.com/kisielk/errcheck"
	@go get -u "github.com/go-swagger/go-swagger/cmd/swagger"

# https://github.com/AlDanial/cloc
cloc:
	@cloc --sdir='Godeps' --not-match-f='Makefile|_test.go' .

# TODO: Get an understanding of how this works
#errcheck:
#	@echo "=== errcheck ==="
#	@errcheck ${REPO}/...

vet:
	@echo "=== go vet ==="
	@go vet ./...

lint:
	@echo "=== go lint ==="
	@golint ./**/*.go

fmt:
	@echo "=== go fmt ==="
	@go fmt ./...

swagger:
	@echo "=== swagger ==="
	@swagger generate spec -o ./swagger-ui/swagger.json

install: test
	@echo "=== go install ==="
	@go install -ldflags=$(GOLDFLAGS)

build: test swagger
	@echo "=== go build ==="
	@mkdir -p bin/
	@go build -ldflags=$(GOLDFLAGS) -o bin/${BINARY_NAME}

dump_schema:
	@echo "=== pg_dump ==="
	@pg_dump -U postgres --schema-only --no-owner mentor_me > schema.sql

test: fmt vet lint errcheck
	@echo "=== go test ==="
	@go test ./... -cover

deploy: test
	# Compile
	@mkdir -p bin/
	GOARCH=amd64 GOOS=linux godep go build -ldflags=$(GOLDFLAGS) -o bin/${BINARY_NAME}
	# Copy binaries
	@scp bin/${BINARY_NAME} $(DEPLOY_HOST):~/
	# Cleanup binaries
	@rm bin/${BINARY_NAME}

clean:
	@rm -rf bin/* pkg/*

.PHONY: setup cloc errcheck vet lint fmt install build test deploy swagger clean dump_schema
