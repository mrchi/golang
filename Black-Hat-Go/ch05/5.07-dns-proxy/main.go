package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/miekg/dns"
)

const PROXY_CONFIG_FILENAME = "proxy.txt"

func parseConfig(filename string) (map[string]string, error) {
	configs := make(map[string]string)
	fh, err := os.Open(filename)
	if err != nil {
		return configs, err
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return configs, fmt.Errorf("%#v is not a valid line", line)
		}
		configs[parts[0]] = parts[1]
	}
	log.Println("configs set to:")
	for k, v := range configs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	return configs, scanner.Err()
}

func main() {
	var configLock sync.RWMutex

	configs, err := parseConfig(PROXY_CONFIG_FILENAME)
	if err != nil {
		log.Fatalln(err)
	}

	dns.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		if len(r.Question) == 0 {
			dns.HandleFailed(w, r)
			return
		}
		// 只处理第一个 Question
		fqdn := r.Question[0].Name

		log.Printf("Receive DNS query for fqdn %s\n", fqdn)

		// 去掉子域
		parts := strings.Split(fqdn, ".")
		if len(parts) > 3 {
			fqdn = strings.Join(parts[len(parts)-3:], ".")
		}

		// 匹配配置表，读锁
		configLock.RLock()
		upstream, ok := configs[fqdn]
		configLock.RUnlock()

		if !ok {
			dns.HandleFailed(w, r)
			return
		}

		// 按配置分流查询
		resp, err := dns.Exchange(r, upstream)
		if err != nil {
			dns.HandleFailed(w, r)
			return
		}

		if err := w.WriteMsg(resp); err != nil {
			dns.HandleFailed(w, r)
			return
		}
	})

	// 开一个单独的 goroutine 监听 SIGUSR1 信号，触发重新读取文件
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGUSR1) // 使用 kill -30 PID

		for sig := range sigs {
			switch sig {
			case syscall.SIGUSR1:
				log.Println("SIGUSR1: reloading records")
				updatedConfigs, err := parseConfig(PROXY_CONFIG_FILENAME)

				if err != nil {
					log.Printf("SIGUSR1: reloading failed, err: %s\n", err)
				} else {
					// 只在读取没问题的情况下才进行更新
					configLock.Lock()
					configs = updatedConfigs
					configLock.Unlock()
				}
			}
		}
	}()

	log.Fatal(dns.ListenAndServe(":53000", "udp", nil))
}
