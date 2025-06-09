.PHONY: test coverage coverage-html clean

PKGS := $(shell go list ./... | grep -v mocks | grep -v docs | grep -v data)
COVERFILE := cover.out

setup:
	go mod tidy && go run data/generate_data.go

run:
	go run cmd/main.go

test:
	go test -coverprofile=$(COVERFILE) $(PKGS)

coverage: test
	go tool cover -func=$(COVERFILE)

coverage-html: test
	go tool cover -html=$(COVERFILE)

clean:
	rm -f $(COVERFILE)
