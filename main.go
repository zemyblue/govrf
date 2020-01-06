package main	// import "github.com/zemyblue/vrf-go"

import "fmt"
import (
	vrf "vrf-go/crypto"
)

func main()  {
	fmt.Println("Hello")
	var pk, sk = vrf.VrfKeygen()
	fmt.Printf("private key: %s, public key: %s\n", pk, sk)
}
