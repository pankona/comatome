all: lint build

build:
	cd $(CURDIR)/cmd/comatome && go build

lint:
	gometalinter --vendor --deadline=300s ./...

clean:
	rm -f $(CURDIR)/cmd/comatome/comatome
	rm -rf $(CURDIR)/vendor
