package diarium

import (
	"net"
	"os"
)

// network stores the app execution data
// related to network connection
type network struct {
	// IP is the machine IP
	IP string `json:"ip"`
	// Hostname is the machine's name
	Hostname string `json:"hostname"`
}

// getNetwork generate a network struct
// with filled fields
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
