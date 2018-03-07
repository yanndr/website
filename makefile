# This how we want to name the binary output
BINARY=website

# These are the values we want to pass for VERSION and BUILD
VERSION=`cat version`
BUILD=`git rev-parse HEAD`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

# Builds the project
build:
	go build ${LDFLAGS} -o ${BINARY}

# Builds the project
linux:
	GOOS=linux go build ${LDFLAGS} -o ${BINARY}

publish:
	gulp webpack
	gulp sass
	gulp vendors
	GOOS=linux go build ${LDFLAGS} -o ${BINARY}
	rm -fr dist
	mkdir dist
	cp website dist/
	cp -r templates dist/templates
	cp -r public dist/public

image: publish
	docker build -t my-website .