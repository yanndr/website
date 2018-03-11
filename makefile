# This how we want to name the binary output
BINARY=./dist/website

# These are the values we want to pass for VERSION and BUILD
VERSION=`cat version`
BUILD=`git rev-parse HEAD`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

# Builds the project
build:dependencies
	go build ${LDFLAGS} -o ${BINARY} ./cmd/website/

# Builds the project
linux:dependencies
	GOOS=linux go build  ${LDFLAGS} -o ${BINARY} ./cmd/website/

dependencies:
	gulp webpack
	gulp sass
	gulp vendors

publish:dependencies
	GOOS=linux go build ${LDFLAGS} -o ${BINARY} ./cmd/website/
	rm -fr dist
	mkdir dist
	cp -r templates dist/templates
	cp -r public dist/public

image: publish
	docker build -t my-website .
test:
	go test -race ./...