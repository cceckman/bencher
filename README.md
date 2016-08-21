# bencher
A little framework for running little benchmarks.

## Why
When working on [`primes`](https://github.com/cceckman/primes) and
[Project Euler](https://projecteuler.net) problems, I got interested in
benchmarking how various every-so-slightly different solutions behave. I wrote
some code that would take a `map[string]func` (where all the functions have the
same signature), wrap it in the
[`testing`](https://golang.org/pkg/testing/) library's
[`Benchmark`](https://golang.org/pkg/testing/#Benchmark) function
at varying levels of intensity, and write out a table with the results.

This generalizes that code.
