#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def solution(s)
  stack = []

  s.split('').each do |bracket|
    if bracket == ')'
      if stack.pop != '('
        return 0
      end
    else
      stack.push(bracket)
    end
  end

  return stack.empty? ? 1 : 0
end

# binding.pry

assert_equal(1, solution("(()(())())"))
assert_equal(1, solution("(()((()))())"))
assert_equal(0, solution("(()()()()"))
assert_equal(0, solution(")"))
assert_equal(0, solution("())"))
assert_equal(0, solution("(()"))
assert_equal(1, solution(""))
