package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/mzky/tls"
)

func main() {

	var ca tls.CACert
	ca.Cert, ca.Key, _ = tls.GenerateRoot()
	ipArray, _ := GetLocalIPList()
	c, k, _ := ca.GenerateServer(ipArray)

	//ca.Cert, _ = tls.ReadRootCertFile("root.cer")
	//ca.Key, _ = tls.ReadPrivKeyFile("root.key")
	//c, k, _ := ca.GenerateServer([]string{"127.0.0.1"})

	//ca.Cert, _ = tls.ReadRootCert([]byte("cert"))
	//ca.Key, _ = tls.ReadPrivKey([]byte("key"))
	//c, k, _ := ca.GenerateServer([]string{"127.0.0.1"})

	_ = tls.WritePEM("server.pem", c)
	_ = tls.WritePEM("server.key", k)
	//fmt.Println(generate.GenerateRoot())
	cert, _ := tls.CertificateInfo("server.pem")
	fmt.Println(cert.NotAfter.Local().Format("2006-01-02_15:04"))
}

func appendIPNet(slice []net.IPNet, element net.IPNet) []net.IPNet {
	if element.IP.IsLinkLocalUnicast() { // ignore link local IPv6 address like "fe80::x"
		return slice
	}

	return append(slice, element)
}

func GetLocalIpNets() (map[string][]net.IPNet, error) {
	iFaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	returnMap := make(map[string][]net.IPNet)
	for _, iFace := range iFaces {
		if iFace.Flags&net.FlagUp == 0 { // Ignore down adapter
			continue
		}

		addrs, err := iFace.Addrs()
		if err != nil {
			continue
		}

		ipNets := make([]net.IPNet, 0)
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPAddr:
				ipNets = appendIPNet(ipNets, net.IPNet{v.IP, v.IP.DefaultMask()})
			case *net.IPNet:
				ipNets = appendIPNet(ipNets, *v)
			}
		}
		returnMap[iFace.Name] = ipNets
	}

	return returnMap, nil
}

func GetLocalIPList() ([]string, error) {
	ipArray := make([]string, 0)
	ipMap, err := GetLocalIpNets()
	if err != nil {
		return nil, err
	}
	mapAddr := make(map[string]string) //去重
	for _, ipNets := range ipMap {
		for _, ipNet := range ipNets {
			mapAddr[ipNet.IP.String()] = ipNet.IP.String()
		}
	}

	for _, ip := range mapAddr {
		ipArray = append(ipArray, strings.TrimSpace(ip))
	}
	return ipArray, nil
}
