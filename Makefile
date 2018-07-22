# Makefile
APPNAME=godoc-example
CURRENT_DIR=`pwd`
PACKAGE_DIRS=`go list -e ./... | egrep -v "binary_output_dir|.git|mocks"`

.PHONY: test test-package-dirs test-report dep mocks dep
test: test-report
test-package-dirs:
	@echo 'Executing unit tests...'
	go vet $(PACKAGE_DIRS)
	go test $(PACKAGE_DIRS) -race -coverprofile=cover.out -covermode=atomic
test-report: test-package-dirs
	@echo 'Generating test coverage report...'
	go tool cover -html=cover.out -o cover.html
dep:
	@dep ensure -update

# docker cmd below
.PHONY:  docker-build docker-run
docker-build:
	docker build . -t $(APPNAME)/v1
docker-run: docker-build
	docker run -it -v $(CURRENT_DIR):/tmp $(APPNAME)/v1