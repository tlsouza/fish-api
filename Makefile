
	go mod tidy

setup-dev: setup
	npm install -g nodemon

run-local:
	go build ./cmd/main.go
	./main

unit-test:
	ENV=testing NEW_RELIC_ENABLED=false go test ./... -cover -v

integration-test:
	export export NEW_RELIC_ENABLED=false && go test ./app/ports/... -cover -v

test: unit-test integration-test

help:
	@echo "Commands:"
	@echo "	check:			Check golang, npm, docker and docker-compose versions"
	@echo "	setup:			Install npm modules(with yarn) to run application"
	@echo "	run-local: 		Build and run the app"
	@echo "	watch-local:		Build and run the app with nodemon"
	@echo "	docker-build:		Build docker image with dockerfile"
	@echo "	run-docker:		Build and run the app with docker"
	@echo "	test:			Run integration and unit tests"
	@echo "	unit-test:		Run unit tests only"
	@echo "	integration-test:	Run integration tests only"
	@echo "	load-test:		Run k6 tests"

check:
	@echo "\n*** Checking versions ***"
	@echo Go: $(shell go version)
	@echo npm: $(shell npm --version)
	@echo Docker: $(shell docker --version) 
	@echo Docker-Compose: $(shell docker-compose --version)
	@echo "\n"
