#!/usr/bin/env bash
export VERSION=0.0.1-dev
export VERSION_GO_VAR=canti/app/version.Version

mkdir ./build
cd build

echo building windows amd64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-X ${VERSION_GO_VAR}=${VERSION} -s -w -extldflags -static -extldflags -static" -o canti.exe canti
upx -8 ./canti.exe
tar -czvf canti_windows_amd64.tar.gz canti.exe
echo - finish

echo building linux amd64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X ${VERSION_GO_VAR}=${VERSION} -s -w -extldflags -static -extldflags -static" -o canti  canti
upx -8 ./canti
tar -czvf canti_linux_amd64.tar.gz canti
echo - finish

echo building windows 386
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags "-X ${VERSION_GO_VAR}=${VERSION} -s -w -extldflags -static -extldflags -static" -o canti.exe canti
upx -8 ./canti.exe
tar -czvf canti_windows_386.tar.gz canti.exe
echo - finish

echo building linux 386
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags "-X ${VERSION_GO_VAR}=${VERSION} -s -w -extldflags -static -extldflags -static" -o canti canti
upx -8 ./canti
tar -czvf canti_linux_386.tar.gz canti
echo - finish
