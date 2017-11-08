# for Library
require 'open-uri'
require 'nokogiri'

# Fetch and parse HTML document
url = 'https://www.ted.com/talks/'

charset = nil
html = open(url) do |file|
	charset = file.charset
	file.read
end

doc = Nokogiri::HTML.parse(html, nil, charset)

# print contents to standard output
puts doc.title
