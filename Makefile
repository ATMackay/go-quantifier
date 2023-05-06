# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

GOBUILD = env GO111MODULE=on go build

test:
	@echo "Running tests..."
	@ go test -v ./...	

volatility:
	@$(GOBUILD) cmd/volatility/main.go 
	@mv main ./build/volatility
	@echo "Run \"./build/volatility --config cmd/volatility/vol.yml\" to launch volatility server."
	@echo "Done building."