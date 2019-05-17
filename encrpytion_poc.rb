require 'openssl'

KEY = "\xBC\x00\xA5\xB5\x0E\xFBN\x1A0\xC7\xC3$\xBA\x00`\xBA"
SIZE = 1312000

def enc(str)
  cipher = OpenSSL::Cipher::AES.new(128, :CBC)
  cipher.encrypt
  cipher.key = KEY
  cipher.padding = 0
  cipher.update(str)
end

def dec(str)
  cipher = OpenSSL::Cipher::AES.new(128, :CBC)
  cipher.decrypt
  cipher.key = KEY
  cipher.padding=0
  cipher.update(str)
end


def run
  alpha = ["a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", " "]

  data = (1..SIZE).map { alpha.sample }.join('')


  puts data.size


  e = enc(data)
  puts "Encrypted size: #{e.size}"

  d = dec(e)
  puts "Decrypted size: #{d.size}"

  puts "Decryption successful: #{d == data}"
end
