package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {
    key := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16};
    block, _ := aes.NewCipher(key);
    plaintext := []byte{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 1, 1, 1, 1, 1, 1};
    iv := []byte{16, 15, 16, 13, 12, 11, 10, 9, 8, 7, 1, 1, 1, 1, 1, 1};
    mode := cipher.NewCBCEncrypter(block, iv)
    dst := make([]byte, 16);
    mode.CryptBlocks(dst, plaintext)
    fmt.Println(len(dst));
}
