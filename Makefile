# build.properties to share with jenkinsfile
include build.properties
APP_BUILD=`git log --pretty=format:'%h' -n 1`
CURRENT_DIR=`pwd`
GO_FLAGS= CGO_ENABLED=0
GO_LDFLAGS= -ldflags="-X main.AppVersion=$(APP_VERSION) -X main.AppName=$(APP_NAME) -X main.AppBuild=$(APP_BUILD)"
GO_BUILD_CMD=$(GO_FLAGS) go build $(GO_LDFLAGS)
DEP_VERSION=0.3.2

vet:
	@go vet

dep:
	@ go get -u github.com/golang/dep/cmd/dep && \
    go get -u github.com/golang/mock/mockgen && \
    go get -u github.com/jstemmer/go-junit-report && \
    go get github.com/axw/gocov/... && \
    go get github.com/AlekSi/gocov-xml && \
    go get -u gopkg.in/alecthomas/gometalinter.v2 && \
	gometalinter.v2 --install

test:
	@go test ./...

test-sonarqube:
	# Run all tests and send to sonarcube
	mkdir -p $(TESTS_DIR) && \
	gocov test ./... | gocov-xml > tests/coverage.xml && \
	go test -v ./... | go-junit-report > tests/test.xml && \
	gometalinter.v2 --vendor --checkstyle > tests/report.xml || true


build: vet test
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GO_BUILD_CMD) -o $(BUILD_DIR)/$(APP_NAME)-linux-amd64

build-docker:
	docker build --build-arg projectName=$(PROJECT_NAME) --build-arg projectOwner=${PROJECT_OWNER} --build-arg appName=$(APP_NAME) -t $(APP_NAME) .

clean:
	rm -Rf $(BUILD_DIR)
