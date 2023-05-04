# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

GOBUILD = env GO111MODULE=on go build

test:
	@echo "Running tests..."
	@ go test -v ./...	

volatility:
	$(GOBIULD) cmd/volatility/main.go
	@mv cmd/indexer ./build
	@echo "Run \"./build/volatility\" to launch volatility server."
	@echo "Done building."