package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	Status int
	Length int
}

type reqBody struct {
	Location string
	Page     int
}

type Response struct {
	TextLines []string `json:"text"`
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
			rec := recover()
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

	startingChar := PageSize * body.Page
	plaintext := bytify(body.Location)

	cipherText := codify(plaintext, Encrypt)
	fmt.Println(cipherText[startingChar : startingChar+50])

	readableText := readable(cipherText[startingChar : startingChar+PageSize])

	response := &Response{
		TextLines: textLines(readableText),
	}

	jsonResponse, _ := json.Marshal(response)
	w.Write(jsonResponse)
}

func codify(plaintext []byte, mode CodecMode) []byte {
	block, _ := aes.NewCipher(key)

	firstPassText := make([]byte, BookSize)
	codec(block, mode).CryptBlocks(firstPassText, plaintext)

	final := make([]byte, BookSize)
	codec(block, mode).CryptBlocks(final, reverse(firstPassText))
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

func readable(text []byte) []string {
	plaintext := make([]string, PageSize)
	for i, v := range text {
		plaintext[i] = NumToCharMap[v]
	}

	return plaintext
}

func textLines(text []string) []string {
	var lines []string

	for i := 0; i < PageSize; i += 80 {
		slc := text[i : i+80]
		lines = append(lines, strings.Join(slc, ""))
	}

	return lines
}

func reverse(slice []byte) []byte {
	rev := make([]byte, len(slice))

	for i, j := len(slice)-1, 0; i >= 0; i, j = i-1, j+1 {
		rev[j] = slice[i]
	}

	return rev
}

func bytify(bookNum string) []byte {
	var num big.Int
	// need to handle error case
	// _, success :=
	num.SetString(bookNum, 10)
	byteSlice := num.Bytes()

	plaintextBytes := make([]byte, BookSize)
	for i, v := range byteSlice {
		plaintextBytes[i] = v
	}

	return plaintextBytes
}
