#!/usr/bin/env ruby

require 'pry'
require 'pp'
require 'test/unit/assertions'

include Test::Unit::Assertions

def upstream?(direction)
  direction == 0
end

def solution(a, b)
  # write your code in Ruby 2.2
  n = a.length

  fishes = []
  survivors = 0

  (0..(n-1)).each do |index|
    current_fish = a[index]
    direction = b[index]

    if upstream?(direction)
      until fishes.empty? do
        downstream_fish = fishes.pop

        if downstream_fish > current_fish
          fishes.push(downstream_fish)
          break
        end
      end
      survivors += 1 if fishes.empty?
    else
      fishes.push(current_fish)
    end
  end

  survivors + fishes.count
end

# binding.pry

assert_equal(2, solution([4, 3, 2, 1, 5], [0, 1, 0, 0, 0]))
assert_equal(1, solution([4], [0]))
assert_equal(2, solution([4, 3], [0, 1]))
assert_equal(1, solution([4, 3], [1, 0]))
assert_equal(1, solution([4, 3, 2], [1, 0, 0]))
assert_equal(2, solution([4, 3, 2, 1], [1, 0, 0, 1]))
