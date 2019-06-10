#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def perimeter(a, b)
  2 * (a + b)
end

def solution(n)
  i = 1
  result = 0

  min = Float::INFINITY

  while i * i < n do
    if n % i == 0

      min = [min, perimeter(i, n/i)].min
    end
    i += 1
  end

  if i * i == n
    min = [min, perimeter(i, i)].min
  end
  min
end

# binding.pry

assert_equal(22, solution(30))
assert_equal(4, solution(1))
assert_equal(10, solution(6))
assert_equal(32, solution(64))
