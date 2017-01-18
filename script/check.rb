require 'open-uri'

GH_RAW_URL = 'https://raw.githubusercontent.com/libharu/libharu/master/include/hpdf.h'

header = open(GH_RAW_URL, &:read)

functions = []

header.gsub(/^#.*$/, '').scan(/HPDF_EXPORT[^;]+;/) do |export|
  func = export.sub(/HPDF_EXPORT\s*\([^\)]+\)/, '').sub(/\([^\)]+\)\s*\;/, '').strip
  functions << func
end

supported = []
tasks = []
functions.sort!
functions.each do |func|
  l = `git grep -l "#{func}"`.strip.lines.length

  tasks << "- [#{l > 0 ? 'x' : ' '}] #{func}"
  if l > 0
    supported << func
  end
end

sPer = (supported.size * 1.0) / (functions.size * 1.0) * 100.0
rev = `git rev-parse --short HEAD 2> /dev/null`
puts "Supported: #{supported.size}/#{functions.size} = #{'%3.2f' % sPer}% @ #{rev}"
puts ""
puts "<details>"
puts "<summary>API Implementation tasks</summary>"
puts ""
puts tasks.join("\n")
puts ""
puts "</details>"
