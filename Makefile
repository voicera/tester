GO ?= go
GOFMT ?= gofmt
GOIMPORTS ?= goimports
ALL_PACKAGES := ./...
LOCAL_PACKAGES := `$(GO) list $(ALL_PACKAGES) | grep -v /vendor/`
LOCAL_PACKAGES_FOLDERS := $(shell find * -maxdepth 0 -type d | grep -v vendor)

all: clean check

check:
	$(GOFMT) -s -w $(LOCAL_PACKAGES_FOLDERS)
	$(GOIMPORTS) -w $(LOCAL_PACKAGES_FOLDERS)
	$(GO) vet $(LOCAL_PACKAGES)
	$(GO) test $(LOCAL_PACKAGES)

clean:
	$(GO) clean $(ALL_PACKAGES)
