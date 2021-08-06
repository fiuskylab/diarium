package diarium

import (
	"net"
	"os"
)

type network struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
}

func getNetwork() network {
	hostname, _ := os.Hostname()
	return network{
		IP:       getIP(),
		Hostname: hostname,
	}
}

// getIP retrieves machine IP address as string
func getIP() string {
	addrs, _ := net.InterfaceAddrs()
	ipAddr := ""
	for _, addr := range addrs {
		if ip, ok := addr.(*net.IPNet); ok &&
			!ip.IP.IsLoopback() &&
			ip.IP.To4() != nil {
			ipAddr = ip.IP.String()
			break
		}
		continue
	}
	return ipAddr
}
