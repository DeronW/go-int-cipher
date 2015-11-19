package int_cipher

import (
	"bytes"
	"crypto/rc4"
	"encoding/binary"
)

func Encrypt8(i int8, k string) int8 {
	out := xor(int2bytes(int(i)), []byte(k), 1)
	return int8(bytes2int(out))
}

func Encrypt16(i int16, k string) int16 {
	out := xor(int2bytes(int(i)), []byte(k), 2)
	return int16(bytes2int(out))
}

func Encrypt32(i int32, k string) int32 {
	out := xor(int2bytes(int(i)), []byte(k), 4)
	return int32(bytes2int(out))
}
func Encrypt64(i int, k string) int64 {
	out := xor(int2bytes(int(i)), []byte(k), 8)
	return int64(bytes2int(out))
}

func Encrypt(i int, k string) int {
	n := Encrypt64(i, k)
	return int(n)
}

func Decrypt8(i int8, k string) int8 {
	return Encrypt8(i, k)
}

func Decrypt16(i int16, k string) int16 {
	return Encrypt16(i, k)
}

func Decrypt32(i int32, k string) int32 {
	return Encrypt32(i, k)
}
func Decrypt64(i int, k string) int64 {
	return Encrypt64(i, k)
}

func Decrypt(i int, k string) int {
	return Encrypt(i, k)
}

func xor(src []byte, k []byte, bits int) []byte {
	s := append(src[:bits], bytes.Repeat([]byte{0}, bits-len(src))...)
	c, err := rc4.NewCipher(k)
	if err != nil {
		panic(err)
	}

	out := make([]byte, len(s))
	c.XORKeyStream(out, s)

	return out
}

func int2bytes(a int) []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, a)
	return buf.Bytes()
}

func bytes2int(b []byte) int {
	//n := 0
	//m := len(b)
	//for i := 0; i < m; i++ {
	//n += int(b[i]) << uint((m-i-1)*8)
	//}
	//return n
	var i int
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.BigEndian, &i)
	return i
}
