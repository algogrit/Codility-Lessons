#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def solution(a)
  n = a.size

  prefix_sum = [0] * (n + 1)

  (1..n).each do |i|
    prefix_sum[i] = a[i - 1] + prefix_sum[i - 1]
  end

  left_idx = min_left_idx = 0
  avg_here = min_avg = ((a[0] + a[1]) / 2.0)

  (2..n-1).each do |i|
    avg_with_prev = (prefix_sum[i + 1] - prefix_sum[left_idx]) / (i - left_idx + 1.0)
    avg_of_two = (a[i - 1] + a[i]) / 2.0

    if avg_of_two < avg_with_prev
      avg_here = avg_of_two
      left_idx = i - 1
    else
      avg_here = avg_with_prev
    end

    if avg_here < min_avg
      min_avg = avg_here
      min_left_idx = left_idx
    end
  end
  min_left_idx
end

# binding.pry

assert_equal(1,
  solution([4, 2, 2, 5, 1, 5, 8]))
