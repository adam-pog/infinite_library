package main

var key = []byte{242, 171, 103, 183, 21, 162, 130, 123, 189, 239, 94, 229, 125, 159, 26, 244}
var iv = []byte{133, 165, 73, 140, 70, 66, 14, 157, 218, 146, 180, 164, 161, 168, 15, 31}

const(
  BookSize = 1312000
  PageSize = 3200
)

type CodecMode string

const(
  Encrypt CodecMode = "encrypt"
  Decrypt CodecMode = "decrypt"
)



// 4 walls of books
// 5 shelves per wall
// 35 books per shelf
