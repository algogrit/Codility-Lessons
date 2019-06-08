#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def solution(a)
  n = a.length

  sums = [0] * n

  (1..n-1).each do |i|
    sums[i] = a[i] - a[i - 1]

    if sums[i-1] > 0
      sums[i] += sums[i-1]
    end
  end

  max_profit = sums.max

  max_profit.nil? ? 0 : sums.max
end

# binding.pry

assert_equal(356, solution([23171, 21011, 21123, 21366, 21013, 21367] ))
assert_equal(0, solution([]))
assert_equal(0, solution([23180, 23171]))
