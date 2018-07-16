package getter

import (
	"fmt"
	"net"
	"os"
        "os/signal"
        "syscall"
)

func FatalCheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func FilterIPV4(ips []net.IP) []string {
	ret := make([]string, 0)
	for _, ip := range ips {
		if ip.To4() != nil {
			ret = append(ret, ip.String())
		}
	}
	return ret
}

func SignalNotify(interrupt chan<- os.Signal) {
        signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
}
