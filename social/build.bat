@echo off
:: Windows
set GOOS=windows
set GOARCH=amd64
go build -o bin/main.exe ./cmd/api

:: Mac (Apple Silicon)
set GOOS=darwin
set GOARCH=arm64
go build -o bin/main ./cmd/api
echo Builds complete! Check the /bin folder.