#!/bin/bash
platforms=("windows/amd64" "windows/386" "darwin/amd64" "linux/amd64")
for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    output_name='makearms-'$GOOS'-'$GOARCH

    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o build/$output_name cmd/makearms/main.go
done