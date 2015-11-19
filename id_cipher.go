package id_cipher

import (
	"bytes"
	"crypto/rc4"
	"encoding/binary"
	"errors"
)

func Encrypt8(i int8, k string) (int8, error) {
	out, err := xor(int2bytes(int(i)), []byte(k), 1)
	return int8(bytes2int(out)), err
}

func Encrypt16(i int16, k string) (int16, error) {
	out, err := xor(int2bytes(int(i)), []byte(k), 2)
	return int16(bytes2int(out)), err
}

func Encrypt32(i int32, k string) (int32, error) {
	out, err := xor(int2bytes(int(i)), []byte(k), 4)
	return int32(bytes2int(out)), err
}
func Encrypt64(i int, k string) (int64, error) {
	out, err := xor(int2bytes(int(i)), []byte(k), 8)
	return int64(bytes2int(out)), err
}

func Encrypt(i int, k string) (int, error) {
	n, err := Encrypt64(i, k)
	return int(n), err
}

func Decrypt8(i int8, k string) (int8, error) {
	return Encrypt8(i, k)
}

func Decrypt16(i int16, k string) (int16, error) {
	return Encrypt16(i, k)
}

func Decrypt32(i int32, k string) (int32, error) {
	return Encrypt32(i, k)
}
func Decrypt64(i int, k string) (int64, error) {
	return Encrypt64(i, k)
}

func Decrypt(i int, k string) (int, error) {
	return Encrypt(i, k)
}

func xor(src []byte, k []byte, bits int) ([]byte, error) {
	if len(k) > 256 {
		return nil, errors.New("key must less than 256")
	}

	s := append(src[:bits], bytes.Repeat([]byte{0}, bits-len(src))...)
	c, _ := rc4.NewCipher(k)
	out := make([]byte, len(s))
	c.XORKeyStream(out, s)

	return out, nil
}

func int2bytes(a int) []byte {
	b := bytes.NewBuffer([]byte{})
	binary.Write(b, binary.LittleEndian, a)
	return b.Bytes()
}

func bytes2int(b []byte) int {
	n := 0
	m := len(b)
	for i := 0; i < m; i++ {
		n += int(b[i]) << uint((m-i-1)*8)
	}
	return n
}
