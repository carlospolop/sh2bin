#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Please indicate the path to the script"
    exit 1
fi

if ! [ -f "$1" ]; then
    echo "The file $1 doesn't exist"
    exit 1
fi

base64 -w0 "$1" > script.sh.b64

goarchList="386 amd64 arm arm64" #ppc64 ppc64le mips mipsle mips64 mips64le mips64p32 mips64p32le ppc s390 s390x sparc sparc64"
goos="linux darwin"

for os in goos; do
    for arch in $goarchList; do
        echo "Building for $arch"
        env GOOS=$os GOARCH=$arch go build -o builds/sh2bin_$os_$arch
        chmod +x builds/sh2bin_$arch
    done
done

rm script.sh.b64
