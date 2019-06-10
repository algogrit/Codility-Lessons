#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def solution(n)
  i = 1
  result = 0
  while i * i < n do
    if n % i == 0
      result += 2
    end
    i += 1
  end
  if i * i == n
    result += 1
  end
  result
end

# binding.pry

# assert_equal(8, solution(24))
assert_equal(8, solution(42))
