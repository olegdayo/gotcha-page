CURRDIR=$(shell pwd)
BINDIR=${CURRDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.51.1
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
PACKAGE=github.com/olegdayo/gotcha-page/backend/cmd/app

.PHONY: all
all: format build test lint

.PHONY: build
build: bindir
	go build -o ${BINDIR}/app ${PACKAGE}

.PHONY: test
test:
	go test ./...

.PHONY: run
run:
	go run ${PACKAGE}

.PHONY: lint
lint: install-lint
	${LINTBIN} run

.PHONY: check
check: format build test lint
	echo "OK"

bindir:
	mkdir -p ${BINDIR}

format: install-smartimports
	${SMARTIMPORTS} -exclude internal/mocks

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

install-smartimports: bindir
	test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})
