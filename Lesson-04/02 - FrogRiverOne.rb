#!/usr/bin/env ruby

require 'pry'

def solution(x, a)
  # write your code in Ruby 2.2
  x_sum = x * (x + 1) / 2

  counter = Array.new(x + 1, 0)

  a.each_with_index do |elem, index|
    next if elem > x
    count = counter[elem]
    x_sum -= elem if count < 1
    counter[elem] = count + 1
    return index if x_sum == 0
  end

  return -1
end

# binding.pry

puts solution(4, [1, 2, 4, 4, 6, 6, 3, 7, 5, 1, 3])
puts solution(3, [1, 2, 4, 4, 6, 6, 3, 7, 5, 1, 3])
puts solution(6, [1, 2, 4, 4, 6, 6, 3, 7, 5, 1, 3])
puts solution(100_000, [1, 2, 100_000, 4, 6, 6, 3, 7, 5, 1, 3])
puts solution(8, [1, 2, 4, 4, 6, 6, 3, 7, 5, 1, 3])
