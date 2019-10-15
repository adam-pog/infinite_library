package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"net/http"
	"strings"
	"time"
	"github.com/julienschmidt/httprouter"
	"log"
	"os"
	"math/big"
	"encoding/json"
)

type statusWriter struct {
	http.ResponseWriter
	Status int
	Length int
}

type reqBody struct {
	Location string
	Page int
}

type Response struct {
  Text string `json:"text"`
}

func (w *statusWriter) WriteHeader(status int) {
	w.Status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.Status == 0 {
		w.Status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.Length += n
	return n, err
}

func main() {
	router := httprouter.New()
	router.POST("/", book)

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: logger(router),
	}

	server.ListenAndServe()
}

func logger(router http.Handler) http.Handler {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Serving http://localhost:8081")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := statusWriter{ResponseWriter: w}

		start := time.Now()
		logger.Printf("%s\t\t%s", r.Method, r.RequestURI)

		defer func() {
			rec := recover();
			if rec != nil {
				sw.WriteHeader(http.StatusInternalServerError)
				fmt.Println(rec)
			}

			logger.Printf(
				"Response\t\t%v\t\t%s\t\t%v\n\n",
				sw.Status,
				r.RemoteAddr,
				time.Since(start),
			)
		}()

		router.ServeHTTP(&sw, r)


	})
}

func book(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
  w.Header().Set("Content-Type", "application/json")


	decoder := json.NewDecoder(r.Body)
	var body reqBody
	//handle err
  decoder.Decode(&body)

	starting_char := PageSize * body.Page
	plaintext := bytify(body.Location)

	cipher_text := codify(plaintext, Encrypt)
	fmt.Println(cipher_text[starting_char:starting_char+50])

  response := &Response{
    Text: readable(cipher_text[starting_char:starting_char+PageSize]),
  }

  json_response, _ := json.Marshal(response)
  w.Write(json_response)
}

func codify(plaintext []byte, mode CodecMode) []byte {
	block, _ := aes.NewCipher(key)

	first_pass_text := make([]byte, BookSize)
	codec(block, mode).CryptBlocks(first_pass_text, plaintext)

	final := make([]byte, BookSize)
	codec(block, mode).CryptBlocks(final, reverse(first_pass_text))
	return final
}

func codec(block cipher.Block, mode CodecMode) cipher.BlockMode {
  if mode == Encrypt {
    return cipher.NewCBCEncrypter(block, iv)
  } else if mode == Decrypt {
    return cipher.NewCBCDecrypter(block, iv)
  }

  return nil
}

func readable(text []byte) string {
	plaintext := make([]string, PageSize)
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

func bytify(bookNum string) []byte {
	var num big.Int
	// need to handle error case
	// _, success :=
	num.SetString(bookNum, 10)
	byteArr := num.Bytes()

	plaintextBytes := make([]byte, BookSize)
	for i, v := range byteArr {
		plaintextBytes[i] = v
	}

	return plaintextBytes
}
