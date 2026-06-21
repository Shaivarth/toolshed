#!/usr/bin/env ruby

def usage
  puts "caesar -- encode or decode text using a Caesar cipher"
  puts ""
  puts "USAGE"
  puts "  ruby caesar.rb encode SHIFT TEXT"
  puts "  ruby caesar.rb decode SHIFT TEXT"
  puts ""
  puts "EXAMPLES"
  puts '  ruby caesar.rb encode 3 "Hello World"'
  puts '  ruby caesar.rb decode 3 "Khoor Zruog"'
  puts ""
  exit 0
end

def caesar(text, shift, decode)
  shift = decode ? -shift : shift

  text.chars.map do |ch|
    if ch.match?(/[a-z]/)
      ((ch.ord - 'a'.ord + shift) % 26 + 'a'.ord).chr
    elsif ch.match?(/[A-Z]/)
      ((ch.ord - 'A'.ord + shift) % 26 + 'A'.ord).chr
    else
      ch
    end
  end.join
end

usage if ARGV.empty? || ARGV[0] == "--help"

unless ARGV.length == 3
  warn "error: expected 3 arguments: <encode|decode> <shift> <text>"
  warn "Run with --help for usage."
  exit 1
end

mode, raw_shift, text = ARGV

unless %w[encode decode].include?(mode)
  warn "error: mode must be 'encode' or 'decode', got '#{mode}'"
  exit 1
end

shift = Integer(raw_shift) rescue nil
unless shift
  warn "error: shift must be an integer, got '#{raw_shift}'"
  exit 1
end

result = caesar(text, shift % 26, mode == "decode")

puts ""
puts "  Mode   : #{mode}"
puts "  Shift  : #{shift}"
puts "  Input  : #{text}"
puts "  Output : #{result}"
puts ""