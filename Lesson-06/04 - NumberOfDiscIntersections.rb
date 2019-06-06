#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def solution(a)
  circle_endpoints = a.each_with_index.flat_map do |radius, j|
    x = j - radius
    y = radius + j

    [[x, 0], [y, 1]]
  end.sort

  intersections = 0
  active_discs = 0
  circle_endpoints.each do |(point, is_end)|
    if is_end == 0
      binding.pry
      intersections += active_discs
      active_discs += 1
    else
      active_discs -= 1
    end

    if intersections > 10E6
      return -1
    end
  end
  intersections
end

# binding.pry

assert_equal(11, solution([1, 5, 2, 1, 4, 0]))
