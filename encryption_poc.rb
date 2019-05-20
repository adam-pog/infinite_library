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
  results = { successful: 0, failures: 0, size_mismatch: 0, num_range_error: 0 }
  alpha = ["a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", " "]

  100.times do
    data = (1..SIZE).map { alpha.sample }.join('')

    e = enc(data)
    e.chars.map do |char|
      num = char.ord
      results[:num_range_error] += 1 if num < 0 || num > 255
    end
    # puts "Encrypted size: #{e.size}"

    d = dec(e)

    d == data ? results[:successful] += 1 : results[:failures] += 1
    results[:size_mismatch] += 1 if d.size != SIZE || e.size != SIZE
    # puts "Decrypted size: #{d.size}"

    # puts "Decryption successful: #{d == data}"
  end

  puts 'Results out of 100: '
  puts "Successful: #{results[:successful]}"
  puts "Failures: #{results[:failures]}"
  puts "Size mismatch: #{results[:size_mismatch]}"
  puts "Num range error: #{results[:num_range_error]}"
end
