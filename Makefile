# ----------------------------------------------------------------------------
# global

.DEFAULT_GOAL = test

# ----------------------------------------------------------------------------
# target

.PHONY: all
all: mod pkg/install

# ----------------------------------------------------------------------------
# include

include hack/make/go.mk

# ----------------------------------------------------------------------------
# overlays

.PHONY: test/gojay
test/gojay: GO_BUILDTAGS+=gojay
test/gojay: TARGET=test/gojay
test/gojay: test

.PHONY: coverage/ci/gojay
coverage/ci/gojay: GO_BUILDTAGS+=gojay
coverage/ci/gojay: coverage/ci
	$(call target)
