# Release specific
VERSION=v0.3.0
CODENAME='Metadata Pride'

# Build parameters
BUILDDIR=build/
BINARYNAME=goca
BUILDTIME=`date +%FT%T%z`
BUILDHASH=`git log -1 --pretty=format:"%h"`
LDFLAGS="\
	-s \
	-w \
	-X \"main.Version=${VERSION}\" \
	-X \"main.Codename=${CODENAME}\" \
	-X \"main.BuildHash=${BUILDHASH}\" \
	-X \"main.BuildTime=${BUILDTIME}\" \
"

ARGS?=

SRC=$(wildcard *.go)

all: run build

build: $(SRC)
	go build -ldflags=$(LDFLAGS) -o ${BUILDDIR}/$(BINARYNAME)_${VERSION} *.go

run: $(SRC)
	go run -ldflags=$(LDFLAGS) *.go $(ARGS)

# test:
# 	GOCA_TEST_SERVER="https://test.goca.io" $(GOTEST) -v -race -cover -count 1 ./...

# test-local:
# 	GOCA_TEST_SERVER="http://localhost:5000" $(GOTEST) -v -race -cover -count 1 ./...

# clean: 
# 	$(GOCLEAN)
# 	rm -rf $(BUILDDIR)

# deps:
# 	$(GOMOD) tidy
# 	$(GOGET) -t -v ./...

# build-linux: $(SRC)
# 	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) ${LDFLAGS} -o $(BINARYNAME)_unix -v $^

.PHONY: all run build 