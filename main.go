package main // import "github.com/zemyblue/govrf"

import (
	"encoding/hex"
	"fmt"
)
import (
	vrf "govrf/crypto"
)

func toHexString(str []byte) string {
	return hex.EncodeToString(str)
}

func main() {
	fmt.Println("Hello")
	var pk, sk = vrf.VrfKeygen()
	fmt.Printf("private key: %s, public key: %s \n", toHexString(pk[:]), toHexString(sk[:]))
}
