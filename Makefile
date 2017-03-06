GO ?= go
ALL_PACKAGES := ./...
LOCAL_PACKAGES := `$(GO) list $(ALL_PACKAGES) | grep -v /vendor/`

all: clean check

check:
	$(GO) fmt $(LOCAL_PACKAGES)
	$(GO) vet $(LOCAL_PACKAGES)
	$(GO) test $(LOCAL_PACKAGES)

clean:
	$(GO) clean $(ALL_PACKAGES)
