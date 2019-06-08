#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def solution(a)
  n = a.length

  (1..n-1).each do |i|
    if a[i-1] > 0
      a[i] += a[i-1]
    end
  end

  a.max
end

# binding.pry

assert_equal(22, solution([3, 4, -3, 2, -10, 3, -1, 3, 3, 3, 4, -3, 2, 3, -1, 3, 3]))
assert_equal(5, solution([3, 2, -6, 4, 0]))
assert_equal(3, solution([3]))
assert_equal(3, solution([3, -1]))
assert_equal(9, solution([3, -1, 4, -2, 5, -3]))
