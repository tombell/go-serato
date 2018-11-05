MODFLAGS=-mod=vendor
TESTFLAGS=-cover

all: test

test:
	go test ${MODFLAGS} ${TESTFLAGS} ./...

.PHONY: all test
