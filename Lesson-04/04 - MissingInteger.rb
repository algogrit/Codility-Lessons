#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def solution(a)
  # write your code in Ruby 2.2
  max = a.max
  counter = Hash.new(0)

  return 1 if max < 1

  a.each do |elem|
    if elem > 0
      counter[elem] += 1
    end
  end

  max.times.each do |key|
    next if key == 0
    return key if counter[key] == 0
  end

  max + 1
end
# binding.pry

assert_equal(5, solution([1, 3, 6, 4, 1, 2]))
assert_equal(4, solution([1, 2, 3]))
assert_equal(1, solution([-1, -3]))
