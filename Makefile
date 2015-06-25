# go-poc Makefile

REPORT_DOCS  		= docs/README.txt
REPORT_COVER  		= app/app.coverprofile
REPORT_COVER_HTML 	= report/app-coverage.html


DOCKER_VERSION ?= latest
DOCKER_IMAGE    = klikindockerhub/gopoc:$(DOCKER_VERSION)


default: build

dep-get:
	go get -v -t ./...

dep-save:
	godep save ./...

build: test
	rm -Rf dist
	mkdir -p dist
	godep go build -o dist/go-poc

install:
	godep go install

dev-run: build
	dist/go-poc

dev-test:
	ginkgo watch -notify -r

dev-coverage: prereport
	ginkgo -r --cover
	godep go tool cover -html=$(REPORT_COVER)

test:
	ginkgo -r

test-docs: prereport
	rm -f $(REPORT_DOCS)
	mkdir -p docs
	ginkgo -r --noColor >> $(REPORT_DOCS)

ci-test: prereport
	ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --cover --trace --race --progress --noColor
	go tool cover -html=$(REPORT_COVER) -o $(REPORT_COVER_HTML)

prereport:
	rm -Rf report
	mkdir -p report

docker-build: test
	docker build --tag $(DOCKER_IMAGE) .

docker-push: docker-build
	docker push $(DOCKER_IMAGE)

.PHONY: test
