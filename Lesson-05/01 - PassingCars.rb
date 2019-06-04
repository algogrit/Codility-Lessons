#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def solution(a)
  # write your code in Ruby 2.2
  # write your code in Ruby 2.2
  count_of_pair = 0
  count_of_east = 0

  a.each do |car|
    if car == 0 # east
      count_of_east += 1
    else
      count_of_pair += count_of_east
      if count_of_pair > 1_000_000_000
        return -1
      end
    end
  end

  count_of_pair
end

# binding.pry

assert_equal(5,
  solution([0, 1, 0, 1, 1]))

assert_equal(0,
  solution([1, 1, 1, 0, 0]))

assert_equal(3,
  solution([1, 0, 1, 0, 1]))

assert_equal(-1,
  solution([0] * 50_000 + [1] * 50_000))
