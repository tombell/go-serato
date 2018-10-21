MODFLAGS=-mod=vendor

all: test

test:
	go test ${MODFLAGS} ./...

.PHONY: all test
