#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def solution(a)
  sorted = a.sort

  negative_dual = sorted.take(2).reduce(:*)

  positive_triplet = sorted.reverse.take(3).reduce(:*)

  negative_triplet = negative_dual * sorted.last

  [negative_triplet, positive_triplet].max
end

# binding.pry

assert_equal(60, solution([-3, 1, 2, -2, 5, 6]))
assert_equal(1_000_000_000_000_000, solution([-3, 100_000, 2, 100_000, 100_000, 6]))
assert_equal(5_000_000, solution([-1, 10, -10, -1000, 500]))
assert_equal(100_000, solution([-1, -10, -10, 1000, 500]))
