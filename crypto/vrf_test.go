package crypto

import "testing"

func TestVrfKeygen(t *testing.T) {
	//var pk, sk = VrfKeygen()
	var pk, sk = "pk", "sk"
	t.Logf("private key: %s, public key: %s", pk, sk)
}
