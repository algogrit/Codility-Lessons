#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

IMPACT_FACTOR = {
  "A" => 1,
  "C" => 2,
  "G" => 3,
  "T" => 4
}

def solution(s, p, q)
  impacts = s.split('').map {|letter| IMPACT_FACTOR[letter]}

  m = p.length

  m.times.map do |index|
    impacts[p[index]..q[index]].min
  end
end

# binding.pry

assert_equal([2, 4, 1],
  solution("CAGCCTA", [2, 5, 0], [4, 5, 6]))
