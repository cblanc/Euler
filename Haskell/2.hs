main = print $ fib 1 1

fib x y = if x + y < 10
					then y
					else fib y x+y