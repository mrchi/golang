package main

import (
	"fmt"
	"log"

	"github.com/miekg/dns"
)

func main() {
	var msg dns.Msg
	fqdn := dns.Fqdn("mrchi.cc")
	msg.SetQuestion(fqdn, dns.TypeA)

	respMsg, err := dns.Exchange(&msg, "8.8.8.8:53")
	if err != nil {
		log.Fatalln(err)
	}

	if len(respMsg.Answer) < 1 {
		log.Fatalln("No records")
	}

	for _, answer := range respMsg.Answer {
		if a, ok := answer.(*dns.A); ok {
			fmt.Println(a.A)
		}
	}
}
