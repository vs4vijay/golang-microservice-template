PROJECT := "golang-microservice-template"
PREFIX  := "/usr/local"

GIT_TAG 	:= `git describe --tags --abbrev=0 --exact-match 2>/dev/null || echo "canary"`
GIT_COMMIT 	:= `git rev-parse HEAD`
BUILD_TIME  := `date -u +"%Y-%m-%dT%H:%M:%SZ"`

LDFLAGS := ""
LDFLAGS += -X=github.com/${PROJECT}/pkg/version.Version=$(GIT_TAG)
LDFLAGS += -X=github.com/${PROJECT}/pkg/version.GitCommit=$(GIT_COMMIT)
LDFLAGS += -X=github.com/${PROJECT}/pkg/version.BuildTime=$(BUILD_TIME)


print-%:
	@echo $* = $($*)

.PHONY: init
init:
	GO111MODULE=on go mod init ${PROJECT}
	GO111MODULE=on cobra init --pkg-name ${PROJECT}

.PHONY: info
info:
	@echo "info..."
	@echo "Version:       		${GIT_TAG}"
	@echo "Git Commit:       	${GIT_COMMIT}"
	@echo "Build Time:       	${BUILD_TIME}"
	@echo ""

.PHONY: format
format:
	goimports -l -w .

.PHONY: lint
lint:
	GO111MODULE=on go mod tidy
	# golint
	# gofmt
	# goimports
	# govet

.PHONY: build
build:
	GO111MODULE=on go build -v -o bin/${PROJECT} -ldflags "${LDFLAGS}" main.go

.PHONY: clean
clean:
	GO111MODULE=on go clean --modcache
	rm -rf bin/*

.PHONY: install
install: build
	cp bin/${PROJECT} ${PREFIX}/bin/${PROJECT}

.PHONY: uninstall
uninstall:
	rm -f ${PREFIX}/bin/${PROJECT}

