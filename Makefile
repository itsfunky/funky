VERSION ?= development

.PHONY: install
install:
	go install -a -ldflags "-X main.Version=${VERSION}" github.com/itsfunky/funky/cmd/funky

.PHONY: test
test:
	go test --tags test -v -cover -covermode=count ./{,aws,cmd,local}/...
