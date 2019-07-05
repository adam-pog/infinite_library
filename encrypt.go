package main

import (
    "crypto/aes"
    "crypto/cipher"
    "fmt"
    "time"
    "strings"
    "net/http"
    "runtime"
    "reflect"
    "io"
    "github.com/julienschmidt/httprouter"
)

func main() {
    mux := httprouter.New()
    mux.GET("/book/:num", book)

    server := http.Server{
      Addr: "localhost:8080",
      Handler: mux,
    }

    fmt.Println("Serving http://localhost:8080")
    server.ListenAndServe()
}

func book(w http.ResponseWriter, r *http.Request, p httprouter.Params){
  num_arr := strings.Split(p.ByName("num"), "")
  plaintext := bytify(num_arr)
  fmt.Println(num_arr)
  fmt.Println(plaintext[0:3199])
  cipher_text := encrypt(key, iv, plaintext)


  // fmt.Println(readable(cipher_text))
  io.WriteString(w, readable(cipher_text[0:3199]))
  io.WriteString(w, "\n\n\n")
  io.WriteString(w, readable(cipher_text[1308800:1312000]))
}

func log(handler http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    name := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
    fmt.Println("Handler function called - " + name)

    start := time.Now()
    handler(w, r)
    fmt.Println(fmt.Sprintf("Time: %s\n", time.Since(start)))
  }
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

func readable(text []byte) string {
    plaintext := make([]string, 3200);
    for i, v := range text {
        plaintext[i] = NumToCharMap[v]
    }

    // fmt.Println(strings.Join(plaintext, ""))
    return strings.Join(plaintext, "")
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
        plaintextBytes[i] = CharToNumMap[v]
    }

    return plaintextBytes
}
