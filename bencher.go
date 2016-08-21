// github.com/cceckman/bencher
// All right reserved? Really? This is the internet and this is source.
// Just give credit where it's due and it's all good.
package bencher

import(
	"flag"
	"fmt"
	"io"
	"os"
	"testing"
	"strings"
)

var(
	outputMode = flag.String("output-mode", "tsv",
		"Output format. Valid values are 'tsv' (columns separated by tabs)," +
		"'csv' (columns separated by commas), or 'col' (columns aligned via the tabwriter package.)")
)

// TODO account for inputs to the function, or explicitly don't.
type Runnable func() fmt.Stringer
type Cases map[string]Runnable

// Run benchmarks on cases, and write the output to f using the specified outputMode.
func Benchmark(cases Cases, output io.Writer, outputMode string) error {
	var sep string
	switch(outputMode) {
	case "tsv","col":
		sep = "\t"
	case "csv":
		sep = ","
	default:
		return fmt.Errorf("Unknown output mode: '%s'", outputMode)
	}
	// Set up writer
	// TODO

	// Write out column headers
	fmt.Fprintln(output, strings.Join([]string{
			"Name",
			"Iterations",
			"Total time (s)",
			"Average time (ns)",
			"Average memory alloced (B)",
			"Average allocation operations",
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
			fmt.Sprint(perfResult.N),
			fmt.Sprint(perfResult.T),
			fmt.Sprint(perfResult.NsPerOp()),
			fmt.Sprintf("%8d", perfResult.AllocedBytesPerOp()),
			fmt.Sprintf("%8d", perfResult.AllocsPerOp()),
		}, sep))
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
