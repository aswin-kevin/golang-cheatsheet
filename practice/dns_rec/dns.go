package dns_rec

import (
	"fmt"
	"net"
)

func DNSrec(){
	ipRec,_ := net.LookupAddr("107.180.100.56")
	for _,ip := range ipRec{
		fmt.Println(ip)
	}
}