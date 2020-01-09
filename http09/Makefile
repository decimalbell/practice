BIN=http09
SRCS=$(wildcard *.go)

.PHONY: show fmt clean

all: $(BIN) 

$(BIN): $(SRCS)
	go build -o $(BIN) $(SRCS)

show:
	@echo $(SRCS)

fmt:
	go fmt

clean:
	rm -vf $(BIN)
