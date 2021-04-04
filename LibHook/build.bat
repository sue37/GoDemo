SET CGO_ENABLED=1 SET GOOS=windows SET GOARCH=amd64 go build -o lib.dll -buildmode=c-shared .\Lib\lib.go
SET CGO_ENABLED=1 SET GOOS=windows SET GOARCH=amd64 go build -o main.exe .\Main\main.go
