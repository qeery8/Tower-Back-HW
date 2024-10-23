package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	countFlag  = flag.Bool("c", false, "count the number of times strings occur")
	dublicFlag = flag.Bool("d", false, "output only duplicate strings")
	uniqueFlag = flag.Bool("u", false, "output only unique strings")
	ignoreCase = flag.Bool("i", false, "ignore case")
	skipFields = flag.Int("f", 0, "skip first fields")
	skipChars  = flag.Int("s", 0, "skip first characters")
)

func normLine(line string, ignoreCase bool, numFields, numChars int) string {
	if ignoreCase {
		line = strings.ToLower(line)
	}

	if numFields > 0 {
		fields := strings.Fields(line)
		if len(fields) > numFields {
			line = strings.Join(fields[numFields:], " ")
		} else {
			line = ""
		}
	}

	if numChars > 0 && len(line) > numChars {
		line = line[numChars:]
	}
	return line
}

func processLines(reader io.Reader, ignoreCase bool, numFields, numChars int) ([]string, map[string]int) {
	scanner := bufio.NewScanner(reader)
	var previosLine string
	counts := make(map[string]int)
	var results []string

	for scanner.Scan() {
		line := scanner.Text()
		normalized := normLine(line, ignoreCase, numFields, numChars)

		if normalized != previosLine {
			results = append(results, line)
			counts[normalized]++
			previosLine = normalized
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read error")
		os.Exit(1)
	}

	return results, counts
}

func printResult(results []string, counts map[string]int, output io.Writer) {
	for _, line := range results {
		normalized := normLine(line, *ignoreCase, *skipFields, *skipChars)
		count := counts[normalized]

		if *countFlag {
			fmt.Fprintf(output, "%d, %s\n", count, line)
		} else if *dublicFlag && count > 1 {
			fmt.Fprintln(output, line)
		} else if *uniqueFlag && count == 1 {
			fmt.Fprintln(output, line)
		} else if !*countFlag && !*dublicFlag && !*uniqueFlag {
			fmt.Fprintln(output, line)
		}
	}
}

func main() {
	flag.Parse()

	if (*countFlag && *dublicFlag) || (*countFlag && *uniqueFlag) || (*dublicFlag && *uniqueFlag) {
		fmt.Fprintln(os.Stderr, "error: flags -c, -d and -u dont use together")
		os.Exit(1)
	}

	var input io.Reader = os.Stdin
	var output io.Writer = os.Stdout

	if flag.NArg() > 0 {
		inputFile, err := os.Open(flag.Arg(0))

		if err != nil {
			fmt.Fprintln(os.Stderr, "error open file input", err)
			os.Exit(1)
		}
		defer inputFile.Close()
		input = inputFile
	}

	if flag.NArg() > 1 {
		outputFile, err := os.Create(flag.Arg(1))

		if err != nil {
			fmt.Fprintln(os.Stderr, "error create file output")
			os.Exit(1)
		}
		defer outputFile.Close()
		output = outputFile
	}

	results, counts := processLines(input, *ignoreCase, *skipFields, *skipChars)

	printResult(results, counts, output)

}
