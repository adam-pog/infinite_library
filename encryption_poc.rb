require 'openssl'
require './mappings'

KEY = "\xBC\x00\xA5\xB5\x0E\xFBN\x1A0\xC7\xC3$\xBA\x00`\xBA"
SIZE = 1312000

def enc(str)
  cipher = OpenSSL::Cipher::AES.new(128, :CBC)
  cipher.encrypt
  cipher.key = KEY
  cipher.padding = 0
  (cipher.update(str).bytes.map{|byte| NUM_TO_CHAR_MAP[byte]}).join('')
end

def dec(str)
  text = str.chars.map{|x| CHAR_TO_NUM_MAP[x].chr }.join('')
  cipher = OpenSSL::Cipher::AES.new(128, :CBC)
  cipher.decrypt
  cipher.key = KEY
  cipher.padding=0
  cipher.update(text)
end


def run
  results = { successful: 0, failures: 0, size_mismatch: 0, num_range_error: 0 }
  alpha = (0..255).to_a.map{|x| NUM_TO_CHAR_MAP[x]}

  1.times do
    data = (1..SIZE).map { alpha.sample }.join('')

    puts data.size
    e = enc(data)
    puts e.size
    e.bytes.map do |num|
      # num = char.ord

      if num < 0 || num > 255
        puts num
        results[:num_range_error] += 1
      end
      # results[:num_range_error] += 1 if num < 0 || num > 255
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
