BIN = certgen

build:
	go build -v -o $(BIN) cmd/main.go

clean:
	rm -f $(BIN)
