#!/usr/bin/env ruby

def usage
  puts "counter -- count lines, words, and characters in a file"
  puts ""
  puts "USAGE"
  puts "  ruby counter.rb FILE"
  puts ""
  puts "EXAMPLES"
  puts "  ruby counter.rb notes.txt"
  puts "  ruby counter.rb caesar.rb"
  puts ""
  exit 0
end

def count(path)
  lines = 0
  words = 0
  chars = 0

  File.foreach(path) do |line|
    lines += 1
    words += line.split.length
    chars += line.length
  end

  { lines: lines, words: words, chars: chars }
end

usage if ARGV.empty? || ARGV[0] == "--help"

unless ARGV.length == 1
  warn "error: expected exactly one file path"
  warn "Run with --help for usage."
  exit 1
end

path = ARGV[0]

unless File.exist?(path)
  warn "error: file not found: #{path}"
  exit 2
end

unless File.readable?(path)
  warn "error: permission denied: #{path}"
  exit 2
end

result = count(path)

puts ""
puts "  File       : #{path}"
puts "  Lines      : #{result[:lines]}"
puts "  Words      : #{result[:words]}"
puts "  Characters : #{result[:chars]}"
puts ""