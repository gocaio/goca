#!/usr/bin/env bash

package=goca/goca.go
package_name=goca

buildDate=$(date -R)
gitTag=$(git describe --tags --abbrev=0)
gitCommit=$(git rev-parse HEAD)
build_dir=build/
ldflags="-X \"main.buildDate=$buildDate\" \
-X \"main.gitTag=$gitTag\" \
-X \"main.gitCommit=$gitCommit\" \
-s \
-w"

platforms=("windows/amd64" "windows/386" "darwin/amd64" "linux/amd64" "linux/386" "linux/arm64" "linux/arm")
go mod download
for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    GOOS=$GOOS GOARCH=$GOARCH go build -o $build_dir$output_name -ldflags "$ldflags" $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    else
        sha256sum $build_dir$output_name > $build_dir$output_name.sha256
    fi
done
