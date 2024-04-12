n = gets.to_i
s = gets.chars
hash = [ ]
n.times do|i|
hash << s[i]
return  p hash.size if hash.include?('A') && hash.include?('B') && hash.include?('C')
end
