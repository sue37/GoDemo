set CGO_ENABLED=1
go build -ldflags="-s -w" -buildmode=c-shared -o lib.dll dll.go