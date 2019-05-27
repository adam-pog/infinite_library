require 'openssl'
require './mappings'

KEY = "\xBC\x00\xA5\xB5\x0E\xFBN\x1A0\xC7\xC3$\xBA\x00`\xBA"
IV = "\xD4\xF1\x95L\xC5\x91\x15&\x8F\x92\xFA\xB0\xE0:T\x8A"
SIZE = 1312000
TEST_ITERATIONS = 10

# SIZE = 32

def enc(str)
  text = str.chars.map{|x| CHAR_TO_NUM_MAP[x].chr }.join('')
  cipher = OpenSSL::Cipher::AES.new(128, :CBC)
  cipher.encrypt
  cipher.key = KEY
  cipher.iv = IV
  cipher.padding = 0

  e = cipher.update(text)

  cipher = OpenSSL::Cipher::AES.new(128, :CBC)
  cipher.encrypt
  cipher.key = KEY
  cipher.iv = IV
  cipher.padding = 0
  (cipher.update(e.reverse).bytes.map{|byte| NUM_TO_CHAR_MAP[byte]}).join('')
end

def dec(str)
  text = str.chars.map{|x| CHAR_TO_NUM_MAP[x].chr }.join('')
  cipher = OpenSSL::Cipher::AES.new(128, :CBC)
  cipher.decrypt
  cipher.key = KEY
  cipher.iv = IV
  cipher.padding=0

  d = cipher.update(text)

  cipher = OpenSSL::Cipher::AES.new(128, :CBC)
  cipher.decrypt
  cipher.key = KEY
  cipher.iv = IV
  cipher.padding=0
  (cipher.update(d.reverse).bytes.map{|byte| NUM_TO_CHAR_MAP[byte]}).join('')
end


def run
  results = { successful: 0, failures: 0, size_mismatch: 0, num_range_error: 0 }
  alpha = (0..255).to_a.map{|x| NUM_TO_CHAR_MAP[x]}

  TEST_ITERATIONS.times do
    data = (1..SIZE).map { alpha.sample }.join('')
    e = enc(data)

    e.bytes.map do |num|
      if num < 0 || num > 255
        puts num
        results[:num_range_error] += 1
      end
    end

    d = dec(e)
    d == data ? results[:successful] += 1 : results[:failures] += 1
    results[:size_mismatch] += 1 if d.size != SIZE || e.size != SIZE
  end

  puts "Results out of #{TEST_ITERATIONS}: "
  puts "Successful: #{results[:successful]}"
  puts "Failures: #{results[:failures]}"
  puts "Size mismatch: #{results[:size_mismatch]}"
  puts "Num range error: #{results[:num_range_error]}"
end
