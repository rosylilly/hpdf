require 'open-uri'

GH_RAW_URL = 'https://raw.githubusercontent.com/libharu/libharu/master/include/hpdf.h'

header = open(GH_RAW_URL, &:read)

functions = []

header.gsub(/^#.*$/, '').scan(/HPDF_EXPORT[^;]+;/) do |export|
  func = export.sub(/HPDF_EXPORT\s*\([^\)]+\)/, '').sub(/\([^\)]+\)\s*\;/, '').strip
  functions << func
end

functions.each do |func|
  l = `git grep -l "#{func}"`.strip.lines.length

  puts "- [#{l > 0 ? 'x' : ' '}] #{func}"
end
