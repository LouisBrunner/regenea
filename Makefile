PACKAGE_PARENT = github.com/LouisBrunner
PACKAGE_NAME   = regenea

TARGETS = cmd/regenea

GO_DEPENDENCIES = \
	github.com/golang/dep/cmd/dep \
	golang.org/x/lint/golint	\
	github.com/jstemmer/go-junit-report	\
	github.com/imsky/junit-merger/...	\
	github.com/modocache/gover \
	github.com/axw/gocov/... \
	github.com/mattn/goveralls \
	github.com/AlekSi/gocov-xml

include ./common.mk
