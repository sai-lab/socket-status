GO_BUILDOPT := -ldflags '-s -w'

imports:
	go get -u golang.org/x/tools/cmd/goimports

dep:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

link:
	mkdir -p $(GOPATH)/src/github.com/sai-lab
	ln -sf $(CURDIR) $(GOPATH)/src/github.com/sai-lab/socket-status
	ln -sf $(CURDIR)/vendor $(CURDIR)/vendor/src

fmt:
	goimports -w *.go lib/*/*.go

build: fmt
	go build $(GO_BUILDOPT) -o bin/socket-status main.go