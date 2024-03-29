package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/miekg/dns"
)

type lookupResult struct {
	IPAddr   string
	Hostname string
}

func main() {
	workerCountFlag := flag.Int("c", 100, "The amount of workers to use.")
	serverFlag := flag.String("server", "8.8.8.8:53", "The DNS server to use.")
	wordListFlag := flag.String("wordlist", "wordlist.txt", "The wordlist to use.")
	flag.Parse()

	domainList := flag.Args()
	if len(domainList) < 1 {
		fmt.Println("Domain is required.")
		os.Exit(1)
	}

	fqdns := make(chan string, *workerCountFlag)
	gather := make(chan []lookupResult)
	tracker := make(chan int)
	var results []lookupResult

	// 读取 wordlist 文件
	fh, err := os.Open(*wordListFlag)
	if err != nil {
		log.Fatalln(err)
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)

	// 创建 workers
	for i := 0; i < *workerCountFlag; i++ {
		go worker(fqdns, gather, tracker, *serverFlag)
	}

	// 分派任务，分派完毕后关闭 fqdns channel
	go func() {
		for scanner.Scan() {
			fqdns <- fmt.Sprintf("%s.%s", scanner.Text(), domainList[0])
		}
		close(fqdns)
	}()

	// 收集结果
	go func() {
		for r := range gather {
			results = append(results, r...)
		}
	}()

	// 通过 tracker 等待所有 worker 都执行完毕后，关闭 gather channel
	for i := 0; i < *workerCountFlag; i++ {
		<-tracker
	}
	close(gather)

	// 输出
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, ' ', 0)
	for _, result := range results {
		fmt.Fprintf(w, "%s\t%s\n", result.Hostname, result.IPAddr)
	}
	w.Flush()
}

func worker(fqdns chan string, gather chan []lookupResult, tracker chan int, serverAddr string) {
	for fqdn := range fqdns {
		log.Printf("Looking up for fqdn %s\n", fqdn)
		results := lookup(fqdn, serverAddr)
		if len(results) > 0 {
			gather <- results
		}
	}
	tracker <- 1
}

func lookupA(fqdn, serverAddr string) ([]string, error) {
	var msg dns.Msg
	var ips []string
	msg.SetQuestion(dns.Fqdn(fqdn), dns.TypeA)
	respMsg, err := dns.Exchange(&msg, serverAddr)
	if err != nil {
		return ips, err
	}
	for _, answer := range respMsg.Answer {
		if a, ok := answer.(*dns.A); ok {
			ips = append(ips, a.A.String())
		}
	}
	return ips, nil
}

func lookupCNAME(fqdn, serverAddr string) ([]string, error) {
	var msg dns.Msg
	var fqdns []string
	msg.SetQuestion(dns.Fqdn(fqdn), dns.TypeCNAME)

	respMsg, err := dns.Exchange(&msg, serverAddr)
	if err != nil {
		return fqdns, err
	}
	for _, answer := range respMsg.Answer {
		if c, ok := answer.(*dns.CNAME); ok {
			fqdns = append(fqdns, c.Target)
		}
	}
	return fqdns, err
}

func lookup(fqdn, serverAddr string) []lookupResult {
	var results []lookupResult
	var cfqdn = fqdn // 用于循环的变量
	for {
		// 一个域名只能有 1 个 CNAME 记录，如果设置了 CNAME 记录，则不能有 A 记录
		// 因此优先查询 CNAME 记录，且只需要取 CNAME 查询结果的第一个
		fqdns, err := lookupCNAME(cfqdn, serverAddr)

		// 不出错且有解析结果的情况下，改为解析该 CNAME 记录的域名
		if err == nil && len(fqdns) > 0 {
			cfqdn = fqdns[0]
			continue
		}

		if err != nil {
			log.Println(err)
		}

		// 其他情况，尝试解析 A 记录
		ips, err := lookupA(cfqdn, serverAddr)
		if err != nil {
			log.Println(err)
			break
		}
		for _, ip := range ips {
			results = append(results, lookupResult{IPAddr: ip, Hostname: fqdn})
		}
		break
	}
	return results
}
