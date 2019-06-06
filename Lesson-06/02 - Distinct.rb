#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def solution(a)
  # write your code in Ruby 2.2
  a.sort.uniq.count
end

# binding.pry

assert_equal(3, solution([2, 1, 1, 2, 3, 1]))
assert_equal(2, solution([2, 1, 2, 1]))
