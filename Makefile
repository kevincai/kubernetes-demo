binary:
	GOOS=$(OS) GOARCH=amd64 go build -o bin/gogreeting -v main.go
