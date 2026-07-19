# Concurrent Health Checker

Concurrently check the health of a URL via its HTTP response code.

## Description and Backstory

This This project explores Go's concurrency through Goroutines and workers. I have always been fascinated by how concurrency works, since it makes a program that performs the same task multiple times not only faster, but also more efficient.
I had also been hearing a lot about Go in general and wanted to see what was going on with the language itself, since it is quite popularly called a 'replacement for Python'. That piqued my interest in the language.project explores Go's concurrency through Go routines
and workers. I have always been fascinated by how concurrency works
since it makes a program that performs same task multiple times not
only faster but also better.

I had also been hearing a lot about Go in general and wanted to
see what was going on with the language itself since it is quite
popularly called as 'replacement of Python'. That piqued my interest
about the language.

## Usage

The program can be used:

```
echo "https://link-here.com" | go run main.go -c 5 -timeout 10
```

Where we piped a link to the program. Here is what the flags mean:

`-c` or `-concurrency` is used to show how many links will be processed
concurrently by the program

`-timeout` is used ot give a timeout of when the program will
stop fetching the failing link.

You can also pipe multiple links by writing them into a file:

```
cat file_with_links.txt | go run main.go
```

## Architecture

The program has a bounded worker pool, i.e. rather than having
as many workers as the number of links passed, we use the concurrency
flag `-c` to determine how many workers will be processing the links
concurrently.

These workers read from a queue rather than the same array therefore
eliminating the risk of any race conditions. The `jobs` channel acts
as the input queue whereas the `results` channel acts as the output queue.

## Hardest Problem Solved

Since this was the first time that I used Go, it took some using to how
everything is declared in this language. Unlike every other language that
I have worked with in the past, Go has the most unique was of declaring
variables.

Understanding how to properly use workers without causing deadlocks was
also one of the hardest things to do. It took not only time but also patience
to understand channels, routines and everything that made concurrency possible
in the program.

## Testing

Mock Table Driven tests have been written for the `Get` function
since that is the only function that could could error in the output.
