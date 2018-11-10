# Required
#  - PACKAGE_NAME: Name of the package (used as output binary name in certain conditions)
#  - PACKAGE_PARENT: Address of the package (e.g. github.com/golang)
#  - TARGETS: List of targets to compile
# Optional
#  - CROOT: Path to the fakeroot (contains system libraries)
#  - DEPENDENCIES: Makefile targets needed to build
#  - TEST_DEPENDENCIES: Makefile targets needed to test
#  - CLEAN_UP_FILES: Exta files to delete on clean
#  - GOFLAGS: flags to pass to Go when compiling/testing/generating
#  - CGO: Override the CGO flags given to Golang
#  - GO_DEPENDENCIES: Go packages needed to build/test/etc
#  - BREW_DEPENDENCIES: system (macOS brew) packages needed to build/test/etc
#  - COVERAGE_RESULTS_FILE: Where to output the coverage file
#  - TEST_RESULTS_FILE: Where to output the unit test results
#  - SIGN_CERT: macOS only, the name of the Keychain certificate used to sign the output binary (needed for `serve`)

# Package
PACKAGE 		 = $(PACKAGE_PARENT)/$(PACKAGE_NAME)

# System
CWD       = $(shell pwd)
SYSTEM_OS = $(shell uname -s)
RM        = rm -f
RM_DIR    = $(RM) -r

# Golang config
OOGOPATH 		 ?= yes # Out-of-GOPATH
GO_PATH       = $(shell go env GOPATH)
GO_SRC_PARENT = $(GO_PATH)/src/$(PACKAGE_PARENT)
GO_SRC_DIR    = $(GO_PATH)/src/$(PACKAGE)

CPATH					= $(if $(findstring yes,$(OOGOPATH)),$(GO_SRC_DIR),$(CWD))
TEST_PROP 		= $(if $(findstring yes,$(OOGOPATH)),ImportPath,Dir)

GOFLAGS 		 ?=
GOEXTRAFLAGS ?=
GOALLFLAGS   ?= $(GOFLAGS) $(GOEXTRAFLAGS)

# Sources
ALL_PACKAGES	     = ./...
TEST_PACKAGES  	   = $(shell go list $(GOALLFLAGS) -f '{{if len .TestGoFiles}}{{.'$(TEST_PROP)'}}{{end}}' $(ALL_PACKAGES))
DEPENDENCIES      ?=
TEST_DEPENDENCIES ?=
CLEAN_UP_FILES    ?=

# Dependencies
GO_DEPENDENCIES ?= none

# Product files
BUILD_DIR				       = .build
COVERAGE_RESULTS_FILE ?= $(BUILD_DIR)/coverage.xml
TEST_RESULTS_FILE     ?= $(BUILD_DIR)/tests.xml
PID    				         = /tmp/$(PACKAGE_NAME).pid

# Package dir
PKGDIR    ?= none
PKGDIR_ARG =
ifneq ($(PKGDIR),none)
	PKGDIR_ARG = -pkgdir $(PKGDIR)
endif

# CGo
CROOT    ?= none
ifeq ($(CROOT),none)
	LLP  =
	CGO ?=
else
	LLP  = LD_LIBRARY_PATH=$(CROOT)/lib
	CGO ?= CGO_CFLAGS="-I$(CROOT)/include $(CGO_CFLAGS)" CGO_LDFLAGS="-L$(CROOT)/lib $(CGO_LDFLAGS)"
endif

# OS dependent
INSTALL_OS_PACKAGES = :
SIGN_PACKAGE = :
POST_BUILD = :

## macOS
ifeq ($(SYSTEM_OS),Darwin)
	SIGN_CERT    ?= Tartaruga
	SIGN_PACKAGE  = codesign -f -s $(SIGN_CERT) $(PACKAGE_NAME)

	BREW_DEPENDENCIES ?= none
ifneq ($(BREW_DEPENDENCIES),none)
	# Fails if already installed
	INSTALL_OS_PACKAGES = brew install $(BREW_DEPENDENCIES) || true
endif

ifneq ($(CROOT),none)
	# Fails if already set
	POST_BUILD = install_name_tool -add_rpath $(CROOT)/lib
endif
endif

################### TARGETS

all: build

dev: build lint test

.PHONY: all dev

################### DEPENDENCIES

get:
	dep ensure -v

update:
	dep ensure -update

.PHONY: get update

################### BUILD

REAL_TARGETS =
define create_target
	$(eval TARGET=$(1))
	$(eval IS_MAIN_PACKAGE=$(filter $(TARGET),$(PACKAGE_NAME)))
	$(eval TARGET_BIN=$(lastword $(subst /, ,$(TARGET))))
	$(eval TARGET_PACKAGE=$(PACKAGE)$(if $(IS_MAIN_PACKAGE),,/$(TARGET)))
	$(eval TARGET_NAME=build_$(TARGET_BIN))
	$(eval REAL_TARGETS += $(TARGET_NAME))

$(TARGET_NAME): $(DEPENDENCIES)
	$(CGO) time go build $(PKGDIR_ARG) $(GOALLFLAGS) -i $(TARGET_PACKAGE)
	@[ -f $(TARGET_BIN) ] && $(POST_BUILD) $(TARGET_BIN) || true
endef

$(foreach target,$(TARGETS),$(eval $(call create_target,$(target))))

build: $(REAL_TARGETS)

generate:
	$(CGO) go generate $(GOALLFLAGS) $(ALL_PACKAGES)

.PHONY: build generate

################### TESTS

CI_TMP_REPORT    = $(BUILD_DIR)/.ci_tmp_report
GOVER_PROFILE    = $(BUILD_DIR)/gover.coverprofile
COV_HTML_FILE    = $(BUILD_DIR)/coverage.html
COV_TEXT_FILE    = $(BUILD_DIR)/coverage.txt

test: tests $(COV_HTML_FILE) show_coverage

ci_test: tests $(TEST_RESULTS_FILE) $(COVERAGE_RESULTS_FILE) coveralls
	@$(RM) $(CI_TMP_REPORT)

coveralls: $(GOVER_PROFILE)
	goveralls -coverprofile=$< -service=travis-ci

as_ci_test:

COVERFILES =
TEST_TARGETS =
define create_test
	$(eval DIR=$(1))
	$(eval IMPORT=$(subst $(CWD),$(PACKAGE),$(DIR)))
	$(eval COVERFILE=$(BUILD_DIR)/$(subst /,.,$(IMPORT)))
	$(eval SUBFOLDER=$(subst $(PACKAGE)/,,$(IMPORT)))
	$(eval TARGET_NAME=test_$(subst /,_,$(SUBFOLDER)))
	$(eval TEST_TARGETS += $(TARGET_NAME))
	$(eval COVERFILES += $(COVERFILE))
	$(eval EXTRA_TEST_ARGS=$(if $(findstring ci_test,$(MAKECMDGOALS)),-v 2>&1 | tee $(CI_TMP_REPORT) | go-junit-report > $(COVERFILE).report.xml; cat $(CI_TMP_REPORT),))

$(TARGET_NAME): $(DEPENDENCIES) $(BUILD_DIR)
	@echo Testing $(TARGET_NAME)
	$(CGO) $(LLP) go test $(GOALLFLAGS) -coverprofile=$(COVERFILE).coverprofile $(IMPORT) $(EXTRA_TEST_ARGS)

$(COVERFILE): $(TARGET_NAME)
endef

$(foreach dir,$(TEST_PACKAGES),$(eval $(call create_test,$(dir))))

tests: $(TEST_DEPENDENCIES) $(TEST_TARGETS)

.PHONY: test ci_test as_ci_test tests

################### REPORTS

$(TEST_RESULTS_FILE): $(BUILD_DIR)
	junit-merger $(shell find $(BUILD_DIR) -name '*report.xml') > $@

$(GOVER_PROFILE): $(COVERFILES) $(BUILD_DIR)
	gover $(BUILD_DIR) $@

$(COV_HTML_FILE): $(GOVER_PROFILE) $(BUILD_DIR)
	go tool cover -html=$< -o $@

$(COV_TEXT_FILE): $(GOVER_PROFILE) $(BUILD_DIR)
	gocov convert $< | gocov report > $@

$(COVERAGE_RESULTS_FILE): $(GOVER_PROFILE) $(BUILD_DIR)
	gocov convert $< | gocov-xml > $@

show_coverage: $(COV_TEXT_FILE)
	@cat $(COV_TEXT_FILE) | tail -n 1

.PHONY: show_coverage

################### LINT

fix_lint:
	@gofmt -w -s $(ALL_PACKAGES)

lint:
	$(CGO) go vet $(ALL_PACKAGES)
	golint -set_exit_status $(ALL_PACKAGES)
	gofmt -d -s .

.PHONY: fix_lint lint

################### CLEAN

clean:
	$(RM) $(PACKAGE_NAME) $(CLEAN_UP_FILES)
	$(RM_DIR) $(BUILD_DIR)

.PHONY: clean

################### SETUP

setup: setup_go setup_deps

setup_deps:
	$(INSTALL_OS_PACKAGES)
ifneq ($(GO_DEPENDENCIES),none)
	go get -u $(GO_DEPENDENCIES)
endif

define create_setup
	$(eval DEP=$(if $(findstring yes,$(OOGOPATH)),setup_link,))

setup_go: $(DEP)
endef

$(eval $(call create_setup))

setup_link: delete_link $(GO_SRC_DIR)

delete_link:
	rm -f $(GO_SRC_DIR)

$(GO_SRC_DIR): $(GO_SRC_PARENT)
ifeq ($(SYSTEM_OS),Darwin)
	ln -s $(CWD) $@
else
	ln -sT $(CWD) $@
endif

show_path:
	@echo $(GO_SRC_DIR)

.PHONY: setup setups_deps setup_go setup_link delete_link show_path

################### DIRECTORIES

$(BUILD_DIR) $(GO_SRC_PARENT):
	@mkdir -p $@

################### HOT RELOAD

serve:
	@make hr_restart ; fswatch -e "docs/.*" -e "$(BUILD_DIR)/.*" -e "$(PACKAGE_NAME)/$(PACKAGE_NAME).*" -e "\.git" -o . | xargs -n1 -I{} make hr_restart || make hr_kill

hr_kill:
	@kill `cat $(PID) 2> /dev/null` 2> /dev/null && rm -f $(PID) || true

hr_warn:
	@echo "------------------------------\n\n\nHot Code Reloading: Rebuilding $(PACKAGE_NAME)\n\n\n------------------------------"

hr_restart: hr_warn hr_kill build
	$(SIGN_PACKAGE)
	$(LLP) ./$(PACKAGE_NAME) & echo $$! > $(PID)

.PHONY: serve hr_kill hr_warn hr_restart
