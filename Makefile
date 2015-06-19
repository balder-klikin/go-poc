# go-poc Makefile

REPORT_XUNIT = report/go-poc-xunit.xml
REPORT_DOCS  = docs/README.md


DOCKER_VERSION ?= latest
DOCKER_IMAGE    = klikindockerhub/go-poc:$(DOCKER_VERSION)


default:
	go build

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
	go test ./...

test-docs: pretest
	rm -f $(REPORT_DOCS)
	mkdir -p docs

pretest:

docker-build: test
	docker build --tag $(DOCKER_IMAGE) .

docker-push: docker-build
	docker push $(DOCKER_IMAGE)

.PHONY: test
