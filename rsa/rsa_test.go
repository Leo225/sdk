package rsa

import "testing"

func TestGenRsaKey(t *testing.T) {
	pub, prv, err := GenRsaKey(2048)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pub)
	t.Log(prv)
}
