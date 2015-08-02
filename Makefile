# See http://peter.bourgon.org/go-in-production/
GO ?= go
CONFIG_FILE = ./conf/local.json
BIN = ./gross
SAMPLE_RSS = https://www.reddit.com/r/golang/.rss

all: clean build

build:
	$(GO) build -o $(BIN)

run: build
	$(BIN) $(SAMPLE_RSS)

run-dev:
	$(GO) run main.go rss.go $(SAMPLE_RSS)

test:
	$(GO) test

clean:
	rm -f $(BIN)
