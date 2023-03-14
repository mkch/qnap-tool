SET GOOS=darwin
SET GOARCH=amd64
go build -o bin/qnap-tool-darwin-amd64

SET GOOS=windows
SET GOARCH=amd64
go build -o bin/qnap-tool-windows-amd64.exe

SET GOOS=linux
SET GOARCH=amd64
go build -o bin/qnap-tool-linux-amd64
