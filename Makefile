gom:
	go get github.com/mattn/gom
	gom -test install

fmt:
	gom exec goimports -w sensu/*.go

test:
	gom test -v -cover -race github.com/hico-horiuchi/ohgibone/sensu

link:
	mkdir -p $(GOPATH)/src/github.com/hico-horiuchi
	if [ ! -d $(GOPATH)/src/github.com/hico-horiuchi/ohgibone ]; then ln -s $(CURDIR) $(GOPATH)/src/github.com/hico-horiuchi/ohgibone; fi

unlink:
	rm $(GOPATH)/src/github.com/hico-horiuchi/ohgibone
	if [ -z "`ls $(GOPATH)/src/github.com/hico-horiuchi`" ]; then rmdir $(GOPATH)/src/github.com/hico-horiuchi; fi
