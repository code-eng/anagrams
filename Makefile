build:
	go build
run:
	./anagrams
restart:
	make build run
test:
	go test -v ./...
