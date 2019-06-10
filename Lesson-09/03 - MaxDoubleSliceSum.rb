#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def solution(a)
  n = a.length

  sum_till_index = [0] * n
  sum_from_index = [0] * n

  (1..n-1).each {|i| sum_till_index[i] = [0, a[i] + sum_till_index[i - 1]].max }

  (0..(n-2)).to_a.reverse.each {|i| sum_from_index[i] = [0, sum_from_index[i + 1] + a[i]].max }

  max = -Float::INFINITY

  (1..(n-2)).each do |i|
    max = [max, sum_from_index[i+1] + sum_till_index[i-1]].max
  end

  max
end

# binding.pry

# assert_equal(17, solution([3, 2, 6, 1, 4, 5, 1, 2]))
assert_equal(17, solution([3, 2, 6, -1, 4, 5, -1, 2]))
