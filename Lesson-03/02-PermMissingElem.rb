#!/usr/bin/env ruby

require 'pry'

def solution(a)
  sum = a.reduce(0, :+)

  np1 = a.length + 1

  sum_of_np1 = (np1 * (np1 + 1)) / 2

  sum_of_np1 - sum
end

binding.pry

puts solution([])
puts solution([1])
puts solution([2])
puts solution([1, 2])
