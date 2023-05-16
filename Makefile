# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

GOBUILD = env GO111MODULE=on go build
GORUN = env GO111MODULE=on go run

all:
	@ make test
	@ make build
	@ make run

test:
	@echo "Running tests..."
	@ go test -cover ./...	

build:
	@$(GOBUILD) cmd/volatility/main.go 
	@mv main ./build/volatility
	@echo "Run \"./build/volatility --config cmd/volatility/vol.yml\" to launch volatility server."
	@echo "Done building."

run:
	@echo "\nRunning volatility service\n"
	@$(GORUN) cmd/volatility/main.go --config cmd/volatility/vol.yml
