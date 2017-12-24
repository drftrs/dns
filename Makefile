default: clean deps dns

clean:
	@rm -f dns

deps:
	go get -u github.com/coredns/coredns
	go get -u github.com/coredns/forward
	go get -u github.com/mholt/caddy
	go get -u github.com/miekg/dns
	go get -u golang.org/x/net/context
	go get -u golang.org/x/text

dns:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -o dns
