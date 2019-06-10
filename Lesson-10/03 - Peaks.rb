#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def peaks(a)
  n = a.length
  peak = []

  (1..n-2).each do |i|
    if a[i] > a[i-1] && a[i] > a[i+1]
      peak << i
    end
  end

  peak
end

def solution(a)
  n = a.length
  peaks = peaks(a)

  count_of_peaks = peaks.count

  return 0 if count_of_peaks == 0

  (1..count_of_peaks).to_a.reverse.each do |block_count|
    next if n % block_count != 0

    block_size = n / block_count
    block_has_peak = [false] * block_count

    peaks.each do |peak|
      block_has_peak[peak / block_size] = true
    end

    return block_count if block_has_peak.all?
  end

  return 0
end

# binding.pry

assert_equal(3, solution([1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2]))
