
build :
	@go build -o dist/nginxlog-exporter .
	@GOARCH=amd64 GOOS=linux go build -o dist/nginxlog-exporter_linux-amd64 .
