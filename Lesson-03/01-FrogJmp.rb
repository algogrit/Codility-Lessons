#!/usr/bin/env ruby

require 'pry'

def solution(x, y, d)
  min_dist = y - x

  min_jumps = min_dist / d

  min_jumps += 1 if min_dist % d != 0

  return min_jumps
end

binding.pry

puts solution(50, 50, 1)
puts solution(50, 70, 1)
puts solution(49, 50, 1)
puts solution(50, 51, 100)
