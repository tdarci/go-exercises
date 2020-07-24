# books

This exercise involves creating a "pipeline", a concept that requires
 channels and goroutines.
 
## Problem 1: Overlapping Words

Given the name of an author, write a program that prints out a list of all
 the words that appear in _all_ of their books (of four letters or more).
 
----------

For example, author SANDY has written three books:

"Baseball Is Fun"
```
This is a book about baseball. It is a story of a time I found myself looking
 for answers, and discovered them in a little, sewn ball.
```

"A Day to Remember"
```
The first time I rode on the bus, I had a ball. There was even an exciting
 moment when a baseball player came on the bus.
```

"Best Sports"
```
Baseball is a sport, but soccer (also known as football) is so much better. I
 have been a soccer player for as long as I can remember. One thing in common
 between these two sports is that they each rely on a ball.
```

If we supply our program with the name "SANDY" we will get back: `baseball
, ball`

----------

To solve this problem, you have been given two functions, shown here:





note: Go piplelines


note: context
