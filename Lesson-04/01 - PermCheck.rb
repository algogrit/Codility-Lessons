#!/usr/bin/env ruby

require 'pry'

def solution(a)
  # write your code in Ruby 2.2
  n = a.length

  counter = Array.new(n, 0)
  a_sum = 0

  a.each do |elem|
    a_sum += elem
    return 0 if elem > n
    current_count = counter[elem] || 0
    return 0 if current_count > 0
    counter[elem] = current_count + 1
  end

  n_sum = n * (n+1) / 2

  return 0 if n_sum != a_sum

  return 1
end

# binding.pry

puts solution([1, 2, 4, 4, 6, 6])
puts solution([2, 3])
puts solution([1, 2, 4, 5])
puts solution([1, 2, 4, 1_000_000_000])
puts solution([1, 2, 4, 3])
