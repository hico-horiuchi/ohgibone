fmt:
	go fmt ./...

test:
	go test ./...

link:
	mkdir -p $(GOPATH)/src/github.com/hico-horiuchi
	ln -s $(CURDIR) $(GOPATH)/src/github.com/hico-horiuchi/ohgibone

unlink:
	rm $(GOPATH)/src/github.com/hico-horiuchi/ohgibone
	rmdir $(GOPATH)/src/github.com/hico-horiuchi
