BINARY := api-demo
VERSION ?= vlatest
PLATFORMS := windows linux darwin
os = $(word 1, $@)

GOVENDOR := $(GOPATH)/bin/govendor
$(GOVENDOR):
	go get -u github.com/kardianos/govendor

$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-$(os)-amd64

vendor: $(GOVENDOR)
	$(GOVENDOR) sync

build: windows linux darwin

clean:
	rm -rf release