# ----------------------------------------------------------------------------
# global

.DEFAULT_GOAL = test

# ----------------------------------------------------------------------------
# target

# ----------------------------------------------------------------------------
# include

include hack/make/go.mk

# ----------------------------------------------------------------------------
# overlays

test/gojay: GO_BUILDTAGS+=gojay
test/gojay: test

coverage/ci/gojay: GO_BUILDTAGS+=gojay
coverage/ci/gojay: coverage/ci
	$(call target)
