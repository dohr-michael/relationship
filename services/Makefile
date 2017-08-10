VERSION=`git rev-parse --short HEAD`

.PHONY: default all docker clean dependencies

default: services

all: clean services

dependencies:
	glide install

services: dependencies
	go build -ldflags "-X github.com/dohr-michael/relationship/services/cfg.Version=${VERSION}"

clean:
	@rm -f services debug
