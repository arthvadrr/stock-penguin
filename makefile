build:
	go build -o bin/stock-penguin

run:
	./bin/stock-penguin

test:
	go test -v ./... -count=1
