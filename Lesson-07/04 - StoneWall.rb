#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def solution(h)
  stones = 0

  peaks = []

  h.each do |height|
    while !peaks.empty? && peaks.last > height do
      peaks.pop
    end
    if !peaks.empty? && peaks.last == height
      next
    else
      stones += 1
      peaks.push(height)
    end
  end

  return stones
end

# binding.pry

assert_equal(7, solution([8, 8, 5, 7, 9, 8, 7, 4, 8]))
assert_equal(2, solution([5, 5, 8, 8]))
assert_equal(2, solution([8, 8, 5, 5]))
assert_equal(3, solution([8, 5, 8]))
assert_equal(2, solution([5, 8, 5]))
assert_equal(3, solution([5, 4, 4, 8]))
assert_equal(4, solution([3, 4, 5, 8]))
assert_equal(3, solution([3, 4, 5, 4]))
