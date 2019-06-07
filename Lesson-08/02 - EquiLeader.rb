#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def is_leader(occurances, count)
  occurances > count / 2.0
end

def leader(a)
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

  return candidate, 0 if !is_leader(count, n)

  return candidate, count
end

def solution(a)
  n = a.length

  candidate, count = leader(a)

  return 0 if count == 0

  occurances = 0
  remaining_occurances = count
  equileader_count = 0

  (0..n-1).each do |index|
    if a[index] == candidate
      occurances += 1
      remaining_occurances -= 1
    end

    if is_leader(occurances, index + 1) && is_leader(remaining_occurances, (n - (index + 1)))
      equileader_count += 1
    end
  end

  equileader_count
end

# binding.pry

assert_equal(5, solution([4, 4, 4, 4, 4, 4]))
assert_equal(3, solution([4, 3, 4, 4, 2, 4]))
assert_equal(3, solution([4, 4, 3, 2, 4, 4]))
assert_equal(2, solution([4, 3, 4, 4, 4, 2]))
assert_equal(0, solution([4, 3, 4, 0, 4, 2]))
