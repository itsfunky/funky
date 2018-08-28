VERSION       ?= development
COVERAGE_FILE ?= coverage.out
TEST_FILES    ?= github.com/itsfunky/funky/{,cmd,internal,providers}/...

ALL_GO_FILES := $(shell find . -type f -name '*.go')

.PHONY           : help
help             : Makefile
	@sed -n 's/^## /	/p' $<

## install       : Compiles and installs funky to your Go bin path.
.PHONY           : install
install          :
	@go install -a -ldflags "-X main.Version=${VERSION}" github.com/itsfunky/funky/cmd/funky

## test          : Runs Go unit tests.
.PHONY           : test
test             :
	@go test --tags test -v $(TEST_FILES)

## cover         : Runs Go unit tests with coverage output.
.PHONY           : cover
cover            : $(COVERAGE_FILE)
$(COVERAGE_FILE) : $(ALL_GO_FILES)
	@go test --tags test -v -cover -coverprofile=$(COVERAGE_FILE) -covermode=count $(TEST_FILES)

## coverhtml     : Runs Go unit tests and opens coverage report in your default browser.
.PHONY           : coverhtml
coverhtml        : $(COVERAGE_FILE)
	@go tool cover -html=$(COVERAGE_FILE)

## clean         : Removes auto-generated files.
.PHONY           : clean
clean            :
	@rm -f $(COVERAGE_FILE)
