package main

import (
	"tlsCert/generate"
)

func main() {
	//generate.GenerateRoot()
	var ca generate.CACert
	ca.Cert, _ = generate.ReadRootCert("root.cer")
	ca.Key, _ = generate.ReadPrivKey("root.key")
	c, k, _ := ca.GenerateServer([]string{"127.0.0.1"})

	generate.WritePEM("server.pem", c)
	generate.WritePEM("server.key", k)
	//fmt.Println(generate.GenerateRoot())
}
