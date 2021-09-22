package main

import (
	"github.com/mzky/tlsCert"
)

func main() {
	//generate.GenerateRoot()
	var ca tlsCert.CACert
	ca.Cert, _ = tlsCert.ReadRootCert("root.cer")
	ca.Key, _ = tlsCert.ReadPrivKey("root.key")
	c, k, _ := ca.GenerateServer([]string{"127.0.0.1"})

	tlsCert.WritePEM("server.pem", c)
	tlsCert.WritePEM("server.key", k)
	//fmt.Println(generate.GenerateRoot())
}
