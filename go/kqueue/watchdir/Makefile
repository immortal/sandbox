.PHONY: all clean

all: clean
	#env GOOS=freebsd GOARCH=arm GOARM=6 go build
	#env GOOS=freebsd GOARCH=amd64 go build
	go build -race

clean:
	go clean -i
	@rm -rf watchpid
