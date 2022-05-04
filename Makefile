build:
	rm -rf out
	mkdir out
	GOOS=darwin GOARCH=amd64 go build -o proto2http cmd/main.go
	tar -czf out/proto2http-darwin-amd64.tar.gz proto2http LICENSE README.md
	rm proto2http
	sha256sum out/proto2http-darwin-amd64.tar.gz
	GOOS=linux GOARCH=386 go build -o proto2http cmd/main.go
	tar -czf out/proto2http-linux-386.tar.gz proto2http LICENSE README.md
	rm proto2http
	sha256sum out/proto2http-linux-386.tar.gz
	GOOS=linux GOARCH=amd64 go build -o proto2http cmd/main.go
	tar -czf out/proto2http-linux-amd64.tar.gz proto2http LICENSE README.md
	rm proto2http
	sha256sum out/proto2http-linux-amd64.tar.gz
	GOOS=linux GOARCH=arm go build -o proto2http cmd/main.go
	tar -czf out/proto2http-linux-arm.tar.gz proto2http LICENSE README.md
	rm proto2http
	sha256sum out/proto2http-linux-arm.tar.gz
	GOOS=linux GOARCH=arm64 go build -o proto2http cmd/main.go
	tar -czf out/proto2http-linux-arm64.tar.gz proto2http LICENSE README.md
	rm proto2http
	sha256sum out/proto2http-linux-arm64.tar.gz
	GOOS=windows GOARCH=386 go build -o proto2http.exe cmd/main.go
	zip out/proto2http-windows-386.zip proto2http.exe LICENSE README.md
	rm proto2http.exe
	sha256sum out/proto2http-windows-386.zip
	GOOS=windows GOARCH=amd64 go build -o proto2http.exe cmd/main.go
	zip out/proto2http-windows-amd64.zip proto2http.exe LICENSE README.md
	rm proto2http.exe
	sha256sum out/proto2http-windows-amd64.zip
	GOOS=windows GOARCH=arm go build -o proto2http.exe cmd/main.go
	zip out/proto2http-windows-arm.zip proto2http.exe LICENSE README.md
	rm proto2http.exe
	sha256sum out/proto2http-windows-arm.zip
