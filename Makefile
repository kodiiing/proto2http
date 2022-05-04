build:
	rm -rf out
	GOOS=darwin GOARCH=amd64 go build -o out/proto2http-darwin-amd64 cmd/main.go
	GOOS=linux GOARCH=386 go build -o out/proto2http-linux-386 cmd/main.go
	GOOS=linux GOARCH=amd64 go build -o out/proto2http-linux-amd64 cmd/main.go
	GOOS=linux GOARCH=arm go build -o out/proto2http-linux-arm cmd/main.go
	GOOS=linux GOARCH=arm64 go build -o out/proto2http-linux-arm64 cmd/main.go
	GOOS=windows GOARCH=386 go build -o out/proto2http-windows-386.exe cmd/main.go
	GOOS=windows GOARCH=amd64 go build -o out/proto2http-windows-amd64.exe cmd/main.go
	GOOS=windows GOARCH=arm go build -o out/proto2http-windows-arm.exe cmd/main.go

