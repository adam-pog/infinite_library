package main

import (
    "crypto/aes"
    "crypto/cipher"
    // "fmt"
    "time"
    "strings"
    "net/http"
    // "runtime"
    // "reflect"
    "io"
    "github.com/julienschmidt/httprouter"
    "log"
    "os"
)

func main() {
    router := httprouter.New()
    router.GET("/book/:num", book)

    server := http.Server{
      Addr: "localhost:8080",
      Handler: logger(router),
    }

    server.ListenAndServe()
}

func logger(router http.Handler) http.Handler {
    logger := log.New(os.Stdout, "http: ", log.LstdFlags)
    logger.Println("Serving http://localhost:8080")

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        router.ServeHTTP(w, r)

        // log request by who(IP address)
        requesterIP := r.RemoteAddr

        logger.Printf("%s\t\t%s", r.Method, r.RequestURI)
        log.Printf(
                "Sent\t\t%s\t\t%v\n\n",
                requesterIP,
                time.Since(start),
        )
    })
}

func book(w http.ResponseWriter, r *http.Request, p httprouter.Params){
  num_arr := strings.Split(p.ByName("num"), "")
  plaintext := bytify(num_arr)
  cipher_text := encrypt(key, iv, plaintext)

  io.WriteString(w, readable(cipher_text[0:3199]))
  io.WriteString(w, "\n\n\n")
  io.WriteString(w, readable(cipher_text[1308800:1312000]))
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
