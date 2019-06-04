#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def solution(s, p, q)
  m = p.length

  m.times.map do |index|
    substring = s[p[index]..q[index]]
    min = 4
    if (substring.include?('A'))
      min = 1
    elsif (substring.include?('C'))
      min = 2
    elsif (substring.include?('G'))
      min = 3
    end
    min
  end
end

# binding.pry

assert_equal([2, 4, 1],
  solution("CAGCCTA", [2, 5, 0], [4, 5, 6]))
