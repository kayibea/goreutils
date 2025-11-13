GO = go
RM = rm -rf
MKDIR = mkdir -p

BIN = goreutils
SRC = main.go $(wildcard cmd/*.go)
CMDS = $(patsubst cmd/%.go,%,$(wildcard cmd/*.go))

.PHONY: all clean link

all: $(BIN)

link: all
	$(MKDIR) bin
	@for cmd in $(CMDS); do \
		ln -sf ../$(BIN) bin/$$cmd; \
	done
	@echo "Activate the environment in your shell by running:"
	@echo "  source ./activate"

$(BIN): $(SRC)
	$(GO) build -o $(BIN) .

clean:
	$(RM) bin/
	$(RM) $(BIN)
