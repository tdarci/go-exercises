# books

This set of exercises involves creating a "pipeline", a concept that requires
 channels and goroutines.
 
## Exercise 1: A Simple Pipeline

For this exercise, you will create a simple Go pipeline.

(The concept of pipelines in Go is explained in [this classic Go blog post on
 the subject](https://blog.golang.org/pipelines). The main points of that
  article are still valid, though some of the specifics may be a little
   dated. The part of the article up to the "Stopping short" section are the
    most relevant.)
    
Head to the file `simple\squares.go`. This file contains a function that
 generates a random number, as well as a function that takes a number and
  squares it. Each function is bit slow at doing its work.
  
The function `NoPipleline()` uses both these functions to generate 50 squares
 and print them out. It also prints the time elapsed. (It runs when you do
  `go run main.go`)
 
**Create a new function, `Pipeline()`, that generates 50 squares concurrently
, by using the pipeline pattern.**

Hint: Your pipeline will run three functions that you write, perhaps named:
* `GenerateNumbers`
* `SquareNumbers`    
* `PrintNumbers`
 
## Exercise 2: Overlapping Words

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
