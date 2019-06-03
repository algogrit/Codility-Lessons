#!/usr/bin/env ruby

require 'pry'
require 'pp'

def solution(n, a)
  # write your code in Ruby 2.2
  counters = Array.new(n, 0)

  a.each do |elem|
    if elem <= n
      counters[elem - 1] += 1
    elsif elem == n + 1
      counters = Array.new(n, counters.max)
    end
  end

  return counters
end

# binding.pry

pp solution(5, [3, 4, 4, 6, 1, 4, 4])
pp solution(5, [1, 2, 4, 4, 6, 6])
pp solution(5, [6, 3])
