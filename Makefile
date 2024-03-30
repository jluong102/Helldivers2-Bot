VERSION="v0.0.0"
BIN="helldivers2-bot"
LD_FLAGS="-s"

build:
	go build -ldflags="${LD_FLAGS}" -o ${BIN} ./cmd/helldivers2-bot/*.go
debug: 
	go build -o ${BIN}-debug ./cmd/helldivers2-bot/*.go
