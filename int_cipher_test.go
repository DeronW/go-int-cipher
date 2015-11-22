package int_cipher

import (
	"testing"
)

var KEY string = "abc"

func Test_int2bytes(t *testing.T) {
	a := int2bytes(1000, 2)
	if a[1] != 0xE8 || a[0] != 0x03 {
		t.Fail()
	}
}

func Test_bytes2int(t *testing.T) {
	a := bytes2int([]byte{0x03, 0xE8})
	if a != 1000 {
		t.Fail()
	}
}

func Test_Encrypt(t *testing.T) {
	a := Encrypt(1, KEY)
	b := Encrypt(3449674024, KEY)
	if a != 3449674024 || b != 1 {
		t.Fail()
	}
}

func Test_Encrypt8(t *testing.T) {
	a := Encrypt8(1, KEY)
	b := Encrypt8(204, KEY)
	if a != 204 || b != 1 {
		t.Fail()
	}
}

func Test_Encrypt16(t *testing.T) {
	a := Encrypt16(1, KEY)
	b := Encrypt16(52636, KEY)
	if a != 52636 || b != 1 {
		t.Fail()
	}
}

func Test_Encrypt32(t *testing.T) {
	a := Encrypt32(1, KEY)
	b := Encrypt32(3449674024, KEY)
	if a != 3449674024 || b != 1 {
		t.Fail()
	}
}

func Test_Encrypt64(t *testing.T) {
	a := Encrypt64(1, KEY)
	b := Encrypt64(14816237120726756722, KEY)
	if a != 14816237120726756722 || b != 1 {
		t.Fail()
	}
}

func Test_Decrypt(t *testing.T) {
	a := Decrypt8(1, KEY)
	b := Decrypt8(204, KEY)

	if a != 204 || b != 1 {
		t.Fail()
	}
}

func Test_Decrypt8(t *testing.T) {
	a := Encrypt8(1, KEY)
	b := Encrypt8(204, KEY)
	if a != 204 || b != 1 {
		t.Fail()
	}
}

func Test_Decrypt16(t *testing.T) {
	a := Encrypt16(1, KEY)
	b := Encrypt16(52636, KEY)
	if a != 52636 || b != 1 {
		t.Fail()
	}
}

func Test_Decrypt32(t *testing.T) {
	a := Encrypt32(1, KEY)
	b := Encrypt32(3449674024, KEY)
	if a != 3449674024 || b != 1 {
		t.Fail()
	}
}

func Test_Decrypt64(t *testing.T) {
	a := Encrypt64(1, KEY)
	b := Encrypt64(14816237120726756722, KEY)
	if a != 14816237120726756722 || b != 1 {
		t.Fail()
	}
}
