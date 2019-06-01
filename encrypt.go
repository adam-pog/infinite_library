package main

import (
    "crypto/aes"
    "crypto/cipher"
    "fmt"
    "math/rand"
    "time"
    "./mappings"
    "strings"
)

const Size = 40 * 80 * 410;

func main() {
    start := time.Now()

    rand.Seed(time.Now().UnixNano());

    key := make([]byte, 16)
    rand.Read(key)

    iv := make([]byte, 16)
    rand.Read(iv)

    textArr := strings.Split(strings.Repeat("Ĥ", Size), "")

    plaintext := bytify(textArr)

    encrypt(key, iv, plaintext)
    enc := encrypt(key, iv, plaintext)

    dec:=decrypt(key, iv, enc)
    readable(dec)

    fmt.Println("Time: ", time.Since(start))
    match := true
    for i, _ := range dec {
        if dec[i] != plaintext[i] {
            match = false;
        }
    }

    fmt.Println("match: ", match)
    fmt.Println("match (final): ", dec[0], dec[1])
    fmt.Println("match (plain): ", plaintext[0], plaintext[1])

    // fmt.Println(dec)
}

func encrypt(key []byte, iv []byte, plaintext []byte) []byte {
    block, _ := aes.NewCipher(key);

    mode := cipher.NewCBCEncrypter(block, iv);
    enc := make([]byte, Size)
    mode.CryptBlocks(enc, plaintext);

    mode = cipher.NewCBCEncrypter(block, iv);
    final := make([]byte, Size)
    mode.CryptBlocks(final, reverse(enc));
    return final
}

func decrypt(key []byte, iv []byte, encryptedText []byte) []byte {
    block, _ := aes.NewCipher(key);

    mode := cipher.NewCBCDecrypter(block, iv);
    dec := make([]byte, Size)
    mode.CryptBlocks(dec, encryptedText);

    mode = cipher.NewCBCDecrypter(block, iv);
    final := make([]byte, Size)
    mode.CryptBlocks(final, reverse(dec));

    return final
}

func readable(text []byte) {
    plaintext := make([]string, Size);
    for i, v := range text {
        plaintext[i] = mappings.NumToCharMap[v]
    }

    // fmt.Println(strings.Join(plaintext, ""))
}

func reverse(arr []byte) []byte {
    rev := make([]byte, len(arr))

    for i, j := len(arr)-1, 0; i >= 0; i, j = i-1, j+1 {
       rev[j] = arr[i]
    }

    return rev
}

func bytify(textArr []string) []byte{
    plaintextBytes := make([]byte, Size);
    for i, v := range textArr {
        plaintextBytes[i] = mappings.CharToNumMap[v]
    }

    return plaintextBytes
}
