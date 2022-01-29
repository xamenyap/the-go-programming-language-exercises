package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

const (
	cmdSHA256 = "sha256"
	cmdSHA384 = "sha384"
	cmdSHA512 = "sha512"
)

func main() {
	cmd := flag.String("cmd", cmdSHA256, "types of hash")
	val := flag.String("value", "", "value to hash")
	flag.Parse()

	switch *cmd {
	case cmdSHA256:
		b := sha256.Sum256([]byte(*val))
		fmt.Fprintf(os.Stdout, "%x\n", b)
	case cmdSHA384:
		b := sha512.Sum384([]byte(*val))
		fmt.Fprintf(os.Stdout, "%x\n", b)
	case cmdSHA512:
		b := sha512.Sum512([]byte(*val))
		fmt.Fprintf(os.Stdout, "%x\n", b)
	default:
		flag.Usage()
	}
}
