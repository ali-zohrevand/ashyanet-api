package services

import (
	"errors"
	"log"
	"net"
	"net/http"
	"strings"
)

var cidrs []*net.IPNet

func init() {
	lancidrs := []string{
		"127.0.0.1/8", "10.0.0.0/8", "169.254.0.0/16", "172.16.0.0/12", "192.168.0.0/16", "::1/128", "fc00::/7",
	}

	cidrs = make([]*net.IPNet, len(lancidrs))

	for i, it := range lancidrs {
		_, cidrnet, err := net.ParseCIDR(it)
		if err != nil {
			log.Fatalf("ParseCIDR error: %v", err) // assuming I did it right above
		}

		cidrs[i] = cidrnet
	}
}

func isLocalAddress(addr string) bool {
	for i := range cidrs {
		myaddr := net.ParseIP(addr)
		if cidrs[i].Contains(myaddr) {
			return true
		}
	}

	return false
}

// Request.RemoteAddress contains port, which we want to remove i.e.:
// "[::1]:58292" => "[::1]"
func ipAddrFromRemoteAddr(s string) string {
	idx := strings.LastIndex(s, ":")
	if idx == -1 {
		return s
	}
	return s[:idx]
}
func GetIpOfRequest(r *http.Request) string {
	hdr := r.Header
	hdrRealIP := hdr.Get("X-Real-Ip")
	hdrForwardedFor := hdr.Get("X-Forwarded-For")

	if len(hdrForwardedFor) == 0 && len(hdrRealIP) == 0 {
		return ipAddrFromRemoteAddr(r.RemoteAddr)
	}

	// X-Forwarded-For is potentially a list of addresses separated with ","
	for _, addr := range strings.Split(hdrForwardedFor, ",") {
		// return first non-local address
		addr = strings.TrimSpace(addr)
		if len(addr) > 0 && !isLocalAddress(addr) {
			return addr
		}
	}

	return hdrRealIP
}
func ExternalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
