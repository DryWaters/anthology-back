build:
	go build -o anthology ./cmd/anthology

run: build
	./anthology

test:
	go test -v ./... -count=1