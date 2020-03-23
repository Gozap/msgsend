BUILD_VERSION    := $(shell cat version)
BUILD_DATE       := $(shell date "+%F %T")
COMMIT_SHA1      := $(shell git rev-parse HEAD)

all: clean
	gox -osarch="darwin/amd64 linux/amd64" \
		-output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}" \
		-ldflags	"-X 'github.com/gozap/msgsend/cmd.version=${BUILD_VERSION}' \
					-X 'github.com/gozap/msgsend/cmd.buildDate=${BUILD_DATE}' \
					-X 'github.com/gozap/msgsend/cmd.commitID=${COMMIT_SHA1}'"

clean:
	rm -rf dist

install:
	go install -ldflags	"-X 'github.com/gozap/msgsend/cmd.version=${BUILD_VERSION}' \
               			-X 'github.com/gozap/msgsend/cmd.buildDate=${BUILD_DATE}' \
               			-X 'github.com/gozap/msgsend/cmd.commitID=${COMMIT_SHA1}'"

.PHONY: all release clean install

.EXPORT_ALL_VARIABLES:

GO111MODULE = on
GOPROXY = https://goproxy.io
GOSUMDB = sum.golang.google.cn
