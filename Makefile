build:
	@GO111MODULE=on go build -ldflags '-s -w' -o durationcheck cmd/main.go

install: build
	@mv durationcheck $(GOPATH)/bin/durationcheck
