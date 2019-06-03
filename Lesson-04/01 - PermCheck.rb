#!/usr/bin/env ruby

require 'pry'

def solution(a)
  # write your code in Ruby 2.2
  seen = {}

  max = a.max
  a.each do |v|
    seen[v] = true
  end

  seen.size == max && max == a.size ? 1 : 0
end

# binding.pry

puts solution([1, 2, 4, 4, 6, 6])
puts solution([2, 3])
puts solution([1, 2, 4, 5])
puts solution([1, 2, 4, 1_000_000_000])
puts solution([1, 2, 4, 3])
