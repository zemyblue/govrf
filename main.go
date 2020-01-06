package main // import "github.com/zemyblue/vrf-go"

import (
	"encoding/hex"
	"fmt"
)
import (
	vrf "vrf-go/crypto"
)

func toHexString(str []byte) string {
	return hex.EncodeToString(str)
}

func main() {
	fmt.Println("Hello")
	var pk, sk = vrf.VrfKeygen()
	fmt.Printf("private key: %s, public key: %s \n", toHexString(pk[:]), toHexString(sk[:]))
}
