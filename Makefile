# go-poc Makefile

REPORT_XUNIT = report/go-poc-xunit.xml
REPORT_DOCS  = docs/README.md


DOCKER_VERSION ?= latest
DOCKER_IMAGE    = klikindockerhub/go-poc:$(DOCKER_VERSION)


default: test

install:
	go install

build:
	rm -Rf dist
	mkdir -p dist
	go build -o dist/go-poc

dev-run:

dev-test:
	goconvey --port 9090

test: pretest
	go test -v

test-docs: pretest
	rm -f $(REPORT_DOCS)
	mkdir -p docs

test-unit: pretest

test-e2e: pretest

test-ci: pretest prereport precoverage

coverage: pretest precoverage

pretest:

prereport:
	rm -Rf report
	mkdir -p report

precoverage:
	rm -Rf coverage

docker-build: test
	docker build --tag $(DOCKER_IMAGE) .

docker-push: docker-build
	docker push $(DOCKER_IMAGE)

.PHONY: test
