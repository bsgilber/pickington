GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

TARGETOS=linux
TARGETARCH=amd64

BUILDDIR=dist
BINARYNAME=lambda
COVERPROFILE=coverage.out

build:
	GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) $(GOBUILD) \
	-o $(BUILDDIR)/$(BINARYNAME) cmd/$(BINARYNAME)/main.go

test:
	mkdir -p $(BUILDDIR)
	$(GOTEST) -coverprofile=$(COVERPROFILE) -outputdir=$(BUILDDIR) -v ./...
	go tool cover -func $(BUILDDIR)/$(COVERPROFILE)

clean:
	rm -rf $(BUILDDIR)
