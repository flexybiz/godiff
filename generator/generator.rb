#!/usr/bin/env ruby
#
# ruby generator.rb <num of strings>
#
require 'securerandom'
require 'ffi'

module GoUuid
  extend FFI::Library
    ffi_lib './gouuid/goUuid.so'
      attach_function :GoUuid, [], :string
end

amount = ARGV[0]
amount ||= 100
p "Generating " + amount.to_s + " lines..."

start = Time.now

#number of different strings
num_diff = rand(3..20)
#array of index of diff strings
arr_num_diff = []
num_diff.times do |i|
  arr_num_diff << rand(1..amount.to_i)
end
p arr_num_diff

out = []
amount.to_i.times do |i|
  #out << SecureRandom.uuid
  out << GoUuid.GoUuid()
end

p out.size
f1 = File.open("../example/first.txt", "w")
f2 = File.open("../example/second.txt", "w")

out.each.with_index do |el, i|
  f1 << el + "\n"
  next if arr_num_diff.include? i
  f2 << el + "\n"
end

f1.close
f2.close

p Time.now - start
