#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def solution(a)
  stack = []

  n = a.length

  a.each do |el|
    if stack.empty? || stack.last == el
      stack.push(el)
    else
      stack.pop
    end
  end

  candidate = stack.first

  count = a.select {|el| el == candidate}.count

  count > n / 2 ? a.index(candidate) : -1
end

# binding.pry

assert_equal(0, solution([3, 4, 3, 2, 3, -1, 3, 3]))
assert_equal(-1, solution([3, 4, 3, 2, 2, -1, 3, 3]))
