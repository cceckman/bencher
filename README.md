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

## How
Take a look at [primebench.go](examples/primebench.go). Basically:

- Close the implementations over their inputs, into closures that return
  `fmt.Stringer`s. (If your function returns builtins, e.g. `bool` or `int`,
  this may require a little type gymnastics.)
- Put the resulting closures in a map, specifically, a
  `map[string]func() fmt.Stringer` (or `bencher.Cases` for short.)
- Run
  [`bencher.AutoBenchmark`](https://godoc.org/github.com/cceckman/bencher#AutoBenchmark)
  with that `map` as the argument.

This will output a tab-separated list of results to `os.Stdout`. If you want
more control, you can use
[`bencher.Benchmark`](https://godoc.org/github.com/cceckman/bencher#Benchmark)
or the `--output-mode` flag to control the output format
(`csv`, the default `tsv`, or `col` for a pre-aligned / human-readable output).
The `Benchmark` method also takes an `io.Writer` rather than just writing to
`os.Stdout`.

