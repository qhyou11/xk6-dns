package dns

import (
	"time"

	"github.com/miekg/dns"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/dns", new(DNS))
}

type DNS struct {
	c dns.Client
}

func (dns *DNS) Setup() {
	dns.c.ReadTimeout = 32 * time.Second
}

func (dns *DNS) Exchange(domain, addr string) (*dns.Msg, time.Duration, error) {
	q := GetRequest(domain)
	return dns.c.Exchange(q, addr)
}

func GetRequest(domain string) *dns.Msg {
	rrType := dns.TypeA
	qclass := uint16(dns.ClassINET)
	var q dns.Question
	q.Name = dns.Fqdn(domain)
	q.Qtype = rrType
	q.Qclass = qclass
	req := &dns.Msg{}
	req.Id = dns.Id()
	req.RecursionDesired = true
	req.Question = []dns.Question{q}
	return req
}
