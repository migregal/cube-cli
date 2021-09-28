.PHONY: all test run clean

all:
	@go build -a

test:
	@go test -v ./...

run:
	go run ./ $(args)

clean:
	rm -f ./cube_cli