BIN = certgen

CMD_SOURCES := $(shell find cmd -name main.go)
TARGETS := $(patsubst cmd/%/main.go,%,$(CMD_SOURCES))

build: $(TARGETS)

clean:
	rm -f $(BIN)

%: cmd/%/main.go
	go build -o $@ $<
