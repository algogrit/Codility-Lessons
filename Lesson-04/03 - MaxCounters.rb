#!/usr/bin/env ruby

require 'pry'
require 'pp'

def solution(n, a)
  # write your code in Ruby 2.2
  current_max = 0
  counters = Hash.new(current_max)

  a.each do |elem|
    if elem <= n
      counters[elem - 1] += 1
    elsif elem == n + 1
      current_max = counters.values.max || current_max
      counters = Hash.new(current_max)
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
