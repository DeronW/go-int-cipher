package int_cipher

import (
	"crypto/rc4"
)

func Encrypt8(i uint8, k string) uint8 {
	out := xor(int2bytes(uint(i), 1), []byte(k))
	return uint8(bytes2int(out))
}

func Encrypt16(i uint16, k string) uint16 {
	out := xor(int2bytes(uint(i), 2), []byte(k))
	return uint16(bytes2int(out))
}

func Encrypt32(i uint32, k string) uint32 {
	out := xor(int2bytes(uint(i), 4), []byte(k))
	return uint32(bytes2int(out))
}
func Encrypt64(i uint, k string) uint64 {
	out := xor(int2bytes(uint(i), 8), []byte(k))
	return uint64(bytes2int(out))
}

func Encrypt(i uint, k string) uint32 {
	return uint32(Encrypt32(uint32(i), k))
}

func Decrypt8(i uint8, k string) uint8 {
	return Encrypt8(i, k)
}

func Decrypt16(i uint16, k string) uint16 {
	return Encrypt16(i, k)
}

func Decrypt32(i uint32, k string) uint32 {
	return Encrypt32(i, k)
}
func Decrypt64(i uint, k string) uint64 {
	return Encrypt64(i, k)
}

func Decrypt(i uint, k string) uint32 {
	return Encrypt(i, k)
}

func xor(src []byte, k []byte) []byte {
	s := src
	c, err := rc4.NewCipher(k)
	if err != nil {
		panic(err)
	}
	out := make([]byte, len(s))
	c.XORKeyStream(out, s)
	return out
}

func int2bytes(a uint, bits int) []byte {
	b := make([]byte, bits)
	for i := 0; i < bits; i++ {
		b[bits-i-1] = byte(a & 0xff)
		a = a >> 8
	}
	return b
}

func bytes2int(b []byte) uint {
	n := uint(0)
	m := len(b)
	for i := 0; i < m; i++ {
		n += uint(b[i]) << (uint(m-i-1) * 8)
	}
	return n
}
