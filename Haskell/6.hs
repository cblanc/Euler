main = print $ squareofsums 100 - sumofsquares 100 
sumofsquares x = sum [a ^ 2 | a <- [1..x]]
squareofsums x = sum [a | a <- [1..x]] ^ 2