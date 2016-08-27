// Package bencher wraps `testing`'s benchmark functionality with output format
// control. It can be used to quickly create benchmark grids.
// See https://github.com/cceckman/bencher for more info.
package bencher

import(
	"flag"
	"fmt"
	"io"
	"os"
	"testing"
	"strings"
	"text/tabwriter"
)

var(
	outputMode = flag.String("output-mode", "tsv",
		"Output format. Valid values are 'tsv' (columns separated by tabs)," +
		"'csv' (columns separated by commas), or 'col' (columns aligned via the tabwriter package.)")
)

// TODO account for inputs to the function, or explicitly don't.
type Runnable func() fmt.Stringer
type Cases map[string]Runnable

// TODO: Provide settings through an interface (/ struct).
// TODO: Allow turnup / turndown of parallelism of benchmark calls.

// Run benchmarks on cases, and write the output to f using the specified outputMode.
func Benchmark(cases Cases, w io.Writer, outputMode string) error {
	// Set up output formatting; tab vs. comma-separated; or buffered and tabwriter-aligned
	var sep string
	var output io.Writer
	switch(outputMode) {
	case "tsv":
		sep = "\t"
		output = w
	case "csv":
		sep = ","
		output = w
	case "col":
		sep = "\t"
		output = tabwriter.NewWriter(w, 0, 0, 1, ' ', tabwriter.AlignRight)
	default:
		return fmt.Errorf("Unknown output mode: '%s'", outputMode)
	}

	// Write out column headers
	fmt.Fprintln(output, strings.Join([]string{
			"Name",
			"Result",
			"Iterations",
			"Total time",
			"Avg time (ns)",
			"Avg memory (bytes)",
			"Avg allocs (ops)",
			"", // Elastic tabstops requires a trailing tab, which strings.Join doesn't provide.
	}, sep))

	for name, function := range cases {
		var evalResult string
		b := func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				evalResult = function().String()
			}
		}
		perfResult := testing.Benchmark(b)

		// Write output
		fmt.Fprintln(output, strings.Join([]string{
			name,
			evalResult,
			fmt.Sprint(perfResult.N),
			fmt.Sprint(perfResult.T),
			fmt.Sprintf("%d", perfResult.NsPerOp()),
			fmt.Sprintf("%d", perfResult.AllocedBytesPerOp()),
			fmt.Sprintf("%d", perfResult.AllocsPerOp()),
			"", // Elastic tabstops requires a trailing tab, which strings.Join doesn't provide.
		}, sep))
	}

	// Flush tabwriter output, if the output channel is a tabwriter.Writer
	switch t := output.(type) {
	case *tabwriter.Writer:
		t.Flush()
	}

	return nil
}

// Run benchmarks on cases, and write the output to stdout using the output mode specified by the command-line flag.
func AutoBenchmark(cases Cases) error {
	// Non-atomic, but I'm not going to worry about that until it breaks.
	if !flag.Parsed() {
		flag.Parse()
	}
	// Call Benchmark with defaults
	return Benchmark(cases, os.Stdout, *outputMode)
}
