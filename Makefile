VERSION ?= development

.PHONY: install
install:
	@go install -a -ldflags "-X main.Version=${VERSION}" github.com/itsfunky/funky/cmd/funky

.PHONY: clean
clean:
	@rm coverage.out

.PHONY: test
test:
	@go test --tags test -v ./{,aws,cmd,local}/...

.PHONY: cover
cover:
	@go test --tags test -v -cover -coverprofile=coverage.out -covermode=count ./{,aws,cmd,local}/...

.PHONY: servecover
servecover: cover
	@go tool cover -html=coverage.out
