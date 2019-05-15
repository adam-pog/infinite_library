require './otp'
require './mappings'

class Babel
  BOOK_SIZE_PROD = 1_322_000
  BOOK_SIZE_DEV = 3

  class << self
    def key
      OTP_KEY # UDPATE KEY TO rand 32?
    end

    def book_from_num(num)
      b10_encrypted_book = num.to_s(32).chars.map{ |x| x.to_i(32) }
      # puts b10_encrypted_book.inspect
      padded_book = pad_zeros(b10_encrypted_book)
      # puts padded_book.inspect

      # padded_book2 = padded_book.map { |x| xor(x, OTP_KEY[(num % 32)]) }
      # puts padded_book2.inspect

      encoded_book = xor_book(padded_book)
      # puts encoded_book.inspect
      map_to_alpha(encoded_book).join('')
    end

    def num_from_book(book)
      encoded_book = map_to_num(book.chars)
       # puts encoded_book.inspect
      encrypted_book = xor_book(encoded_book)
       # puts encrypted_book.inspect
      b32_book = encrypted_book.map{ |x| x.to_s(32) }
       # puts b32_book.inspect
      b32_book.join('').reverse.to_i(32)
    end

    private

    def pad_zeros(arr)
      arr.insert(0, *[0] * (BOOK_SIZE_DEV - arr.size))
    end

    def xor_book(book)
      book.map.with_index do |x, i|
        # puts "#{x} xor #{OTP_KEY[i]} = #{xor(x, OTP_KEY[i])}"
        x ^ OTP_KEY[i]
      end
    end

    def map_to_alpha(book)
      book.map{ |x| NUM_MAP[x] }
    end

    def map_to_num(book)
      book.map{ |x| ALPHA_MAP[x] }
    end

    # def xor(a, b)
    #   [a ^ b, 28].min
    # end
  end
end
