package main

import (
	"log"
	"net"

	"github.com/miekg/dns"
)

func main() {
	dns.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		var resp dns.Msg
		resp.SetReply(r)
		for _, q := range r.Question {
			a := dns.A{
				Hdr: dns.RR_Header{
					Name:   q.Name,
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
					Ttl:    0,
				},
				A: net.ParseIP("127.0.0.1").To4(),
			}
			resp.Answer = append(resp.Answer, &a)
		}
		w.WriteMsg(&resp)
	})
	// dig @localhost -p 53000 google.com
	log.Fatal(dns.ListenAndServe(":53000", "udp", nil))
}
