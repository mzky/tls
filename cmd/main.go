package main

import (
	"github.com/mzky/tls"
)

func main() {

	tls.GenerateRoot()

	var ca tls.CACert
	ca.Cert, _ = tls.ReadRootCert("root.cer")
	ca.Key, _ = tls.ReadPrivKey("root.key")
	c, k, _ := ca.GenerateServer([]string{"127.0.0.1"})

	tls.WritePEM("server.pem", c)
	tls.WritePEM("server.key", k)
	//fmt.Println(generate.GenerateRoot())
}
