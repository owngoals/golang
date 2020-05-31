package golang

import "testing"

func TestHashHmac256(t *testing.T) {
	data := "123456"
	secret := "secret_2020"

	s := "ef0895d66f01e3e8affa57ce4b8345c2a305f33b829314420d40f225e2da97db"
	h := HashHmac256(data, secret)

	if s != h {
		t.FailNow()
	}
}
