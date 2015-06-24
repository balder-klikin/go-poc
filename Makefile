# go-poc Makefile

REPORT_XUNIT = report/go-poc-xunit.xml
REPORT_DOCS  = docs/README.md


DOCKER_VERSION ?= latest
DOCKER_IMAGE    = klikindockerhub/go-poc:$(DOCKER_VERSION)


default:
	go build

get:
	godep get

install:
	go install

build:
	rm -Rf dist
	mkdir -p dist
	go build -o dist/go-poc

dev-run: build
	dist/go-poc

dev-test:
	ginkgo watch -notify -r

test: pretest
	ginkgo -r

test-docs: pretest
	rm -f $(REPORT_DOCS)
	mkdir -p docs

ci-test: pretest
	ginkgo -r --noColor

pretest:

docker-build: test
	docker build --tag $(DOCKER_IMAGE) .

docker-push: docker-build
	docker push $(DOCKER_IMAGE)

.PHONY: test
