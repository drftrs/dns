package drftrs

import (
	"log"
	"strings"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
	"golang.org/x/net/context"
)

// DNS struct to implement the plugin.Handler interface.
type DNS struct {
	Next      plugin.Handler
	Blacklist []string
}

// ServeDNS implements the plugin.Handler interface.
func (d *DNS) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}
	if state.QType() == dns.TypeA && d.IsBlocked(r) {
		return dns.RcodeRefused, nil
	}
	return plugin.NextOrFailure(d.Name(), d.Next, ctx, w, r)
}

// Name implements the Handler interface.
func (d *DNS) Name() string { return "drftrs" }

func (d *DNS) IsBlocked(r *dns.Msg) bool {
	for _, q := range r.Question {
		domain := strings.TrimSpace(q.Name[:len(q.Name)-1])
		for _, blacklisted := range d.Blacklist {
			if strings.HasSuffix(domain, blacklisted) {
				log.Printf("Matched black list: %s\n", domain)
				return true
			}
		}
	}
	return false
}
