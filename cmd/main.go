package main

import (
	"github.com/mzky/tls"
)

func main() {

	var ca tls.CACert
	ca.Cert, ca.Key, _ = tls.GenerateRoot()
	c, k, _ := ca.GenerateServer([]string{"127.0.0.1"})

	ca.Cert, _ = tls.ReadRootCertFile("root.cer")
	ca.Key, _ = tls.ReadPrivKeyFile("root.key")
	//c, k, _ := ca.GenerateServer([]string{"127.0.0.1"})

	//ca.Cert, _ = tls.ReadRootCert([]byte("cert"))
	//ca.Key, _ = tls.ReadPrivKey([]byte("key"))
	//c, k, _ := ca.GenerateServer([]string{"127.0.0.1"})

	_ = tls.WritePEM("server.pem", c)
	_ = tls.WritePEM("server.key", k)
	//fmt.Println(generate.GenerateRoot())
}
