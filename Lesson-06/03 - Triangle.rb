#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def solution(a)
  # write your code in Ruby 2.2
  n = a.length
  sorted = a.sort

  return 0 if a.length < 3

  for i in (0..n-3) do
    p = sorted[i]
    q = sorted[1+i]
    r = sorted[2+i]

    is_p_small = p < (q + r)
    is_q_small = q < (r + p)
    is_r_small = r < (q + p)

    return 1 if is_p_small && is_q_small && is_r_small
  end

  return 0
end

# binding.pry

assert_equal(1, solution([10, 2, 5, 1, 8, 20]))
assert_equal(0, solution([10, 50, 5, 1]))
