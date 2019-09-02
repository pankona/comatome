all: lint build test

install:
	cd $(CURDIR)/cmd/comatome && go install

build:
	cd $(CURDIR)/cmd/comatome && go build

test:
	go test -cover ./...

lint:
	golangci-lint run --deadline 300s ./...

clean:
	rm -f $(CURDIR)/cmd/comatome/comatome
