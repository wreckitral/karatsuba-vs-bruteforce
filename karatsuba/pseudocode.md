# Karatsuba Algoritm

Procedure karatsuba(x: int64, y: int64) -> int64
Declare:
max_digits: int64
positive: boolean
x_high: int64
x_low: int64
y_high: int64
y_low: int64
z0: int64
z1: int64
z2: int64

Begin:
max_digits ← 0
positive ← true

if x = 0 OR y = 0 then
    return 0
end if

if (x > 0 AND y < 0) OR (x < 0 AND y > 0) then
    positive ← false
end if

if x < 10 OR y < 10 then
    return x * y
end if

x_digits ← getDecimalDigits(x)
y_digits ← getDecimalDigits(y)

if x_digits >= y_digits then
    max_digits ← x_digits / 2
else
    max_digits ← y_digits / 2
end if

x_high, x_low ← getHighAndLowDigits(x, max_digits)
y_high, y_low ← getHighAndLowDigits(y, max_digits)

z0 ← karatsuba(x_low, y_low)
z1 ← karatsuba((x_low + x_high), (y_low + y_high))
z2 ← karatsuba(x_high, y_high)

if positive then
    return (z2 * 10^(2 * max_digits)) + (z1 - z2 - z0) * 10^max_digits + z0
else
    return -((z2 * 10^(2 * max_digits)) + (z1 - z2 - z0) * 10^max_digits + z0)
end if



Procedure getDecimalDigits(num: int64) -> int64
Declare:
result: int64

Begin:
result ← 0

arduino

if num = 0 then
    return 1
end if

if num < 0 then
    num ← -num
end if

while num > 0 do
    result ← result + 1
    num ← num / 10
end while

return result

End procedure



Procedure getHighAndLowDigits(num: int64, digits: int64) -> (int64, int64)
Declare:
divisor: int64

Begin:
divisor ← 10^digits

if num >= divisor then
    return (num / divisor, num % divisor)
else
    return (0, num)
end if

