build: main.go
	go build -o build
main: 
	go run ./main.go
.PHONY: clean
clean:
