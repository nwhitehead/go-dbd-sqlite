include $(GOROOT)/src/Make.inc

TARG=github.com/thomaslee/go-dbd-sqlite
GOFILES=$(shell find . -name '*.go' -not -name '*test*.go')

include $(GOROOT)/src/Make.pkg

