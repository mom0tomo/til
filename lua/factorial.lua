#!/usr/bin/env lua
-- 標準入力から得た数字を階乗する
function fact(n)
	if n == 0 then
		return 1
	else
		return n * fact(n - 1)
	end
end

print("enter a number: ")
a = io.read("*number")
print(fact(a))
