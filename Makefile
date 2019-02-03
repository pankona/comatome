
all: lint build

build:
	go build

lint:
	go mod vendor
	GO111MODULE=off gometalinter --vendor --deadline=300s ./...

clean:
	rm -f $(CURDIR)/github-contribution-checker
	rm -rf $(CURDIR)/vendor
