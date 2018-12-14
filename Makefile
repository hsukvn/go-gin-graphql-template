.PHONY: all setup build run clean
.DEFAULT_GOAL=all

GRAPHQL_SCHEMA_BINDATA := graphql/schema/bindata.go

all: build

setup:
	go get -u github.com/jteeuwen/go-bindata/...

build: $(GRAPHQL_SCHEMA_BINDATA)
	go build

run: $(GRAPHQL_SCHEMA_BINDATA)
	go run main.go

clean:
	rm -f graphql/schema/bindata.go
	go clean

$(GRAPHQL_SCHEMA_BINDATA):
	go-bindata -ignore=\.go -pkg=schema -o=graphql/schema/bindata.go graphql/schema/...


