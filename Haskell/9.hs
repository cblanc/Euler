-- A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

-- a2 + b2 = c2
-- For example, 32 + 42 = 9 + 16 = 25 = 52.

-- There exists exactly one Pythagorean triplet for which a + b + c = 1000.
-- Find the product abc.

main = print $ ptuple $ take 1 [(a,b,c) | c <- squares, b <- [x | x<-squares, x < c, 1000 - sqrt x - sqrt c > 0], a <- [x | x<-squares, x < b], a + b == c, sqrt a + sqrt b + sqrt c == 1000]
squares = [ x^2 | x <- [1..1000] ]
ptuple [(a, b, c)] = sqrt a * sqrt b * sqrt c