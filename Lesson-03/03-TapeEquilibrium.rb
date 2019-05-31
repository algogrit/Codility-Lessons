#!/usr/bin/env ruby

require 'pry'

def solution(a)
  # write your code in Ruby 2.2
  n = a.length
  remaining_a = a[1..-1]

  left_sum = a[0]
  right_sum = remaining_a.reduce(0, :+)

  min_difference = (left_sum - right_sum).abs

  remaining_a.each_with_index do |elem, index|
    left_sum += elem
    right_sum -= elem

    next if index == n - 2

    difference = (left_sum - right_sum).abs
    min_difference = [min_difference, difference].min
  end

  min_difference
end

binding.pry

puts solution([3, 1, 2, 3, 4])
puts solution([-1000, 1000])
puts solution([1000, 1000])
puts solution([-10, -10, 20])
