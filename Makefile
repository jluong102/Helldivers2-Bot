BIN="helldivers2-bot"

build:
	go build -o ${BIN} ./cmd/helldivers2-bot/*.go
debug: 
	go build -o ${BIN}-debug ./cmd/helldivers2-bot/*.go
