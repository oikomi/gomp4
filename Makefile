all: build

build:
	mkdir -p bin
	go build -o gomp4  gomp4.go

clean:
	rm -rf bin
