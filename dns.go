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

<<<<<<<<<<<<<<  ✨ Codeium Command ⭐  >>>>>>>>>>>>>>>>
// SetupWithTimeout sets the timeout for DNS lookups. The timeout is specified
// in seconds.
<<<<<<<  3e11df6d-fa8a-4e7d-96ea-fa062450fa2c  >>>>>>>
func (dns *DNS) SetupWithTimeout(timeout_time time.Duration) {
	dns.c.ReadTimeout = timeout_time * time.Second
}

func (dns *DNS) Setup() {
	dns.c.ReadTimeout = 5 * time.Second
}

func (dns *DNS) Exchange(domain, addr string) (string, time.Duration) {
	start := time.Now()
	q := GetRequest(domain)
	msg, _, err := dns.c.Exchange(q, addr)
	elapsed := time.Since(start)
	if err != nil {
		return "", elapsed
	}
	return msg.String(), elapsed
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
