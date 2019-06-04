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
  n = s.length
  m = p.length

  prefix_sums = {}

  IMPACT_FACTOR.each do |letter, _|
    prefix_sums[letter] = Array.new(n + 1, 0)
    (1..n).each do |i|
      adder = s[i - 1] == letter ? 1 : 0
      prefix_sums[letter][i] = prefix_sums[letter][i - 1] + adder
    end
  end

  p.zip(q).map do |from, to|
    IMPACT_FACTOR.map do |letter, impact|
      if prefix_sums[letter][from] < prefix_sums[letter][to + 1]
        impact
      end
    end.reject(&:nil?).min
  end
end

# binding.pry

assert_equal([2, 4, 1],
  solution("CAGCCTA", [2, 5, 0], [4, 5, 6]))
