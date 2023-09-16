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
	fmt.Printf("%#v\n", msg)
	fmt.Printf("%#v\n", msg.Question)

	respMsg, err := dns.Exchange(&msg, "8.8.8.8:53")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%#v\n", respMsg)
	fmt.Printf("%#v\n", respMsg.Answer)
}
