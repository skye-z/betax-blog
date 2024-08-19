#!/usr/bin/env bash

echo "Start packaging..."

go mod download
go mod tidy

mkdir ./out
cp LICENSE ./out/LICENSE

generate(){
    CGO_ENABLED=0 GOOS=$1 GOARCH=$2 go build -o betax-blog -ldflags '-s -w'
    mv betax-blog ./out/betax-blog-$2-$1
    cd out
    tar -zcf betax-blog-$2-$1.tar.gz betax-blog-$2-$1 LICENSE
    rm -rf ./betax-blog-$2-$1
    cd ../
}

echo "[1] Linux from amd64"
generate linux amd64
echo "[2] Linux from arm64"
generate linux arm64