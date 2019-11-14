GOC=go
GO111MODULE=on

PKG_SOURCES=$(wildcard pkg/*/*.go)
DIRECTORIES=$(wildcard pkg/*/)
MOCKS=$(foreach x, $(DIRECTORIES), mocks/$(x))


.PHONY: build test test_race lint vet install-deps coverage mocks clean-mocks

all: genesis

genesis: | install-deps
	$(GOC) build ./...

test:
	go test ./...

test_race:
	go test ./... -race 

lint:
	golint $(go list ./... | grep -v mocks)

vet:
	go vet $(go list ./... | grep -v mocks)

install-deps:
	go get ./...

clean-mocks:
	rm -rf mocks

mocks: $(MOCKS)
	
$(MOCKS): mocks/% : %
	mockery -output=$@ -dir=$^ -all
	
#install-mock:
#	go get github.com/golang/mock/gomock
#	go install github.com/golang/mock/mockgen
