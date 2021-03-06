SRCROOT := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))

COVEROUT=$(SRCROOT)/cover.out

AT = cd $(SRCROOT)
RM = rm -fv
GOBUILD = go build -v
GOFMTCHECK = test -z `gofmt -l -s *.go | tee /dev/stderr`
GOTEST = go test -v
GOCOVER = $(GOTEST) -race -coverprofile=/tmp/cover.out -covermode=atomic && cat /tmp/cover.out >> $(COVEROUT)

.PHONY: all
all: fmt build cover

.PHONY: bootstrap
bootstrap:
	glide install --strip-vendor

.PHONY: clean
clean: clean-cover clean-pepcli clean-papcli clean-pdpserver clean-egen

.PHONY: fmt
fmt: fmt-pdp fmt-pdp-yast fmt-pdp-jcon fmt-pdpctrl-client fmt-papcli fmt-pep fmt-pepcli fmt-pepcli-requests fmt-pepcli-test fmt-pepcli-perf fmt-pdpserver fmt-plugin fmt-egen

.PHONY: build
build: build-pepcli build-papcli build-pdpserver build-plugin build-egen

.PHONY: test
test: test-pdp test-pdp-yast test-pdp-jcon test-pep test-plugin

.PHONY: cover
cover: cover-pdp cover-pdp-yast cover-pdp-jcon cover-pep cover-plugin

$(COVEROUT):
	echo > $(COVEROUT)

# Per package clean targets
.PHONY: clean-cover
clean-cover:
	@$(RM) $(COVEROUT)

.PHONY: clean-pepcli
clean-pepcli:
	@$(AT)/pepcli && $(RM) pepcli

.PHONY: clean-papcli
clean-papcli:
	@$(AT)/papcli && $(RM) papcli

.PHONY: clean-pdpserver
clean-pdpserver:
	@$(AT)/pdpserver && $(RM) pdpserver

.PHONY: clean-egen
clean-egen:
	@$(AT)/egen && $(RM) egen

# Per package format targets
.PHONY: fmt-pdp
fmt-pdp:
	@echo "Checking PDP format..."
	@$(AT)/pdp && $(GOFMTCHECK)

.PHONY: fmt-pdp-yast
fmt-pdp-yast:
	@echo "Checking PDP YAST format..."
	@$(AT)/pdp/yast && $(GOFMTCHECK)

.PHONY: fmt-pdp-jcon
fmt-pdp-jcon:
	@echo "Checking PDP JCon format..."
	@$(AT)/pdp/jcon && $(GOFMTCHECK)

.PHONY: fmt-pdpctrl-client
fmt-pdpctrl-client:
	@echo "Checking PDP control client library format..."
	@$(AT)/pdpctrl-client && $(GOFMTCHECK)

.PHONY: fmt-papcli
fmt-papcli:
	@echo "Checking PAP CLI format..."
	@$(AT)/papcli && $(GOFMTCHECK)

.PHONY: fmt-pep
fmt-pep:
	@echo "Checking PEP client library format..."
	@$(AT)/pep && $(GOFMTCHECK)

.PHONY: fmt-pepcli
fmt-pepcli:
	@echo "Checking PEP CLI format..."
	@$(AT)/pepcli && $(GOFMTCHECK)

.PHONY: fmt-pepcli-requests
fmt-pepcli-requests:
	@echo "Checking PEP CLI requests package format..."
	@$(AT)/pepcli/requests && $(GOFMTCHECK)

.PHONY: fmt-pepcli-test
fmt-pepcli-test:
	@echo "Checking PEP CLI test package format..."
	@$(AT)/pepcli/test && $(GOFMTCHECK)

.PHONY: fmt-pepcli-perf
fmt-pepcli-perf:
	@echo "Checking PEP CLI perf package format..."
	@$(AT)/pepcli/perf && $(GOFMTCHECK)

.PHONY: fmt-pdpserver
fmt-pdpserver:
	@echo "Checking PDP server format..."
	@$(AT)/pdpserver && $(GOFMTCHECK)

.PHONY: fmt-plugin
fmt-plugin:
	@echo "Checking PE-CoreDNS Middleware format..."
	@$(AT)/contrib/coredns/policy && $(GOFMTCHECK)
	@$(AT)/contrib/coredns/policy/policytap && $(GOFMTCHECK)

.PHONY: fmt-egen
fmt-egen:
	@echo "Checking EGen format..."
	@$(AT)/egen && $(GOFMTCHECK)

# Per package build targets
.PHONY: build-pepcli
build-pepcli:
	$(AT)/pepcli && $(GOBUILD)

.PHONY: build-papcli
build-papcli:
	$(AT)/papcli && $(GOBUILD)

.PHONY: build-pdpserver
build-pdpserver:
	$(AT)/pdpserver && $(GOBUILD)

.PHONY: build-plugin
build-plugin:
	$(AT)/contrib/coredns/policy && $(GOBUILD)
	$(AT)/contrib/coredns/policy/policytap && $(GOBUILD)

.PHONY: build-egen
build-egen:
	$(AT)/egen && $(GOBUILD)

.PHONY: test-pdp
test-pdp:
	$(AT)/pdp && $(GOTEST)

.PHONY: test-pdp-yast
test-pdp-yast:
	$(AT)/pdp/yast && $(GOTEST)

.PHONY: test-pdp-jcon
test-pdp-jcon:
	$(AT)/pdp/jcon && $(GOTEST)

.PHONY: test-pep
test-pep:
	$(AT)/pep && $(GOTEST)

.PHONY: test-plugin
test-plugin:
	$(AT)/contrib/coredns/policy && $(GOTEST)
	$(AT)/contrib/coredns/policy/policytap && $(GOTEST)

# Per package test coverage targets
.PHONY: cover-pdp
cover-pdp: | $(COVEROUT)
	$(AT)/pdp && $(GOCOVER)

.PHONY: cover-pdp-yast
cover-pdp-yast: | $(COVEROUT)
	$(AT)/pdp/yast && $(GOCOVER)

.PHONY: cover-pdp-jcon
cover-pdp-jcon: | $(COVEROUT)
	$(AT)/pdp/jcon && $(GOCOVER)

.PHONY: cover-pep
cover-pep: | $(COVEROUT)
	$(AT)/pep && $(GOCOVER)

.PHONY: cover-plugin
cover-plugin: | $(COVEROUT)
	$(AT)/contrib/coredns/policy && $(GOCOVER)
	$(AT)/contrib/coredns/policy/policytap && $(GOCOVER)
