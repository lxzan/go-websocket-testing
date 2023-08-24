PWD=$(shell pwd)
GOOS=linux
GOARCH=amd64
CGO_ENABLED=0

clean:
	rm ./bin/*

build:
	CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -o bin/gws-${GOOS}-${GOARCH} cmd/gws/main.go
	CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -o bin/gorilla-${GOOS}-${GOARCH} cmd/gorilla/main.go
	CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -o bin/gobwas-${GOOS}-${GOARCH} cmd/gobwas/main.go
	CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -o bin/nhooyr-${GOOS}-${GOARCH} cmd/nhooyr/main.go
	CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -o bin/nbio-${GOOS}-${GOARCH} cmd/nbio/main.go

