package main

import (
	"net"
	"strings"
	"sync"

	"github.com/droundy/goopt"
)

var server = goopt.String([]string{"-s", "--server"}, "0.0.0.0", "Server to bind")
var port = goopt.String([]string{"-p", "--port"}, "52500", "Port to bind")

func main() {
	goopt.Parse(nil)
	/* Lets prepare a address at any address at port 10001*/
	ports := strings.Split(*port, ",")
	var wg sync.WaitGroup
	for _, uport := range ports {
		wg.Add(1)
		go func(port string) {
			ServerAddr, err := net.ResolveUDPAddr("udp", *server+":"+port)
			if err != nil {
				wg.Done()
				return
			}

			/* Now listen at selected port */
			ServerConn, err := net.ListenUDP("udp", ServerAddr)
			if err != nil {
				wg.Done()
				return
			}
			defer ServerConn.Close()
			buf := make([]byte, 4)

			for {
				for i := range buf {
					buf[i] = 0
				}
				_, addr, err := ServerConn.ReadFromUDP(buf)
				if err != nil {
					continue
				}
				ServerConn.WriteToUDP(buf,addr)
			}
			wg.Done()
		}(uport)
	}
	wg.Wait()
}
