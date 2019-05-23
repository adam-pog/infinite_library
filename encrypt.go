package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
  "math/rand"
  "time"
)

func main() {
    rand.Seed(time.Now().UnixNano());
    key := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16};
    block, _ := aes.NewCipher(key);
    size := 40 * 80 * 410;
    plaintext := make([]byte, size);
    rand.Read(plaintext)

    iv := []byte{16, 15, 16, 13, 12, 11, 10, 9, 8, 7, 1, 1, 1, 1, 1, 1};
    mode := cipher.NewCBCEncrypter(block, iv);
    dst := make([]byte, size);
    mode.CryptBlocks(dst, plaintext);
    fmt.Println(len(dst));
    fmt.Println(dst[0]);
    fmt.Println(plaintext[0]);

//    s1 := rand.NewSource(time.Now().UnixNano())
//    r1 := rand.New(s1)


}
