SET GOOS=linux
SET GOARCH=amd64
SET CGO_ENABLED=0
SET GOPATH=%GOPATH%;%cd%
SET GOBIN=%cd%/bin
go build -o daping  main.go

PAUSE

EXIT