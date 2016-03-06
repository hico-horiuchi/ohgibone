gom:
	go get github.com/mattn/gom
	gom -test install

fmt:
	gom exec goimports -w sensu/*.go

test:
	gom test -v -cover -race ./...

link:
	mkdir -p $(GOPATH)/src/github.com/hico-horiuchi
	ln -s $(CURDIR) $(GOPATH)/src/github.com/hico-horiuchi/ohgibone

unlink:
	rm $(GOPATH)/src/github.com/hico-horiuchi/ohgibone
	rmdir $(GOPATH)/src/github.com/hico-horiuchi
