def fizz_buzz(n)
  if n % 15 == 0
    puts 'Fizz Buzz'
  elsif
    n % 3 == 0
    puts 'Fizz'
  elsif n % 5 == 0
    puts 'Buzz'
  else
    puts n.to_s
  end
end

fizz_buzz(4)
fizz_buzz(100)
fizz_buzz(30)
