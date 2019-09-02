all: lint build

install:
	cd $(CURDIR)/cmd/comatome && go install

build:
	cd $(CURDIR)/cmd/comatome && go build

lint:
	golangci-lint run --deadline 300s ./...

clean:
	rm -f $(CURDIR)/cmd/comatome/comatome
