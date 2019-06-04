#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def solution(a, b, k)
  # write your code in Ruby 2.2
  n_in_a = -1
  n_in_b = 0

  include_a = a % k == 0 ? 1 : 0
  n_in_a = a / k - include_a if a > 0

  n_in_b = b / k if b > 0

  n_in_b - n_in_a
end

# binding.pry

assert_equal(3, solution(6, 11, 2))
assert_equal(5, solution(6, 14, 2))
assert_equal(4, solution(7, 14, 2))
assert_equal(12, solution(0, 11, 1))
assert_equal(2, solution(0, 2_000_000_000, 2_000_000_000))
assert_equal(1, solution(2_000_000_000, 2_000_000_000, 2_000_000_000))
