package main

import (
	"fmt"
	"github.com/bernd/certgen"
)

func main() {
	config := certgen.CertConfig{
		CertFile: "/tmp/yolo.crt",
		CertOrganization: "Example, Inc.",
		KeyFile: "/tmp/yolo.key",
		KeyBits: 2048,
		DnsNames: []string{"localhost", "127.0.0.1"},
	}

	if err := certgen.GenerateCert(config); err != nil {
		fmt.Println("ERROR:", err)
	}
}
