#!/usr/bin/env ruby

require 'pry'
require 'pp'

def solution(n, a)
  # write your code in Ruby 2.2
  max_value = 0
  counters = Hash.new(max_value)

  a.each do |elem|
    if elem <= n
      counters[elem - 1] += 1
      if max_value < counters[elem - 1]
        max_value = counters[elem - 1]
      end
    elsif elem == n + 1
      counters = Hash.new(max_value)
    end
  end

  n.times.map do |i|
    counters[i]
  end
end

# binding.pry

pp solution(5, [3, 4, 4, 6, 1, 4, 4])
pp solution(5, [1, 2, 4, 4, 6, 6])
pp solution(5, [6, 3])
