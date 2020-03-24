package crypto

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func toString(s []uint8) string {
	return hex.EncodeToString(s)
}

func TestVrfKeygen(t *testing.T) {
	var pk, sk = VrfKeygen()
	//var pk, sk = "pk", "sk"
	t.Logf("private key: %s, public key: %s", toString(sk[:]), toString(pk[:]))
}

func TestVrfPrivkey_Prove(t *testing.T) {
	var pk, sk = VrfKeygen()
	var message = []byte("This is test message.")
	t.Logf("private key: %s, public key: %s", toString(sk[:]), toString(pk[:]))
	var proof, ok = sk.Prove(message)
	t.Logf("proof: %s, ok: %t", toString(proof[:]), ok)

	var output, ok2 = proof.Hash()
	t.Logf("output : %s, ok: %t", toString(output[:]), ok2)

	var ok3, output2 = pk.Verify(proof, message)
	t.Logf("output2: %s, ok: %t", toString(output2[:]), ok3)

	if !bytes.Equal(output[:], output2[:]) {
		t.Error("output not matches")
	}
}
