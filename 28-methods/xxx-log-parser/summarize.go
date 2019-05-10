// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/inancgumus/learngo/28-methods/xxx-log-parser/report"
)

// summarize prints the parsing results.
//
// it prints the errors and returns if there are any.
//
// --json flag encodes to json and prints.
func summarize(sum *report.Summary, errors ...error) {
	if errs(errors...) {
		return
	}

	if args := os.Args[1:]; len(args) == 1 && args[0] == "--json" {
		encode(sum)
		return
	}
	stdout(sum)
}

// encodes the summary to json
func encode(sum *report.Summary) {
	out, err := json.MarshalIndent(sum, "", "\t")
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(out)
}

// prints the summary to standard out
func stdout(sum *report.Summary) {
	const format = "%-30s %10s %20s\n"
	const formatValue = "%-30s %10d %20d\n"

	fmt.Printf(format, "DOMAIN", "VISITS", "TIME SPENT")
	fmt.Println(strings.Repeat("-", 65))

	next, cur := sum.Iterator()
	for next() {
		rec := cur()
		fmt.Printf(formatValue, rec.Domain, rec.Visits, rec.TimeSpent)
	}

	fmt.Printf("\n"+formatValue, "TOTAL",
		sum.Total().Visits, sum.Total().TimeSpent,
	)
}

// this variadic func simplifies the multiple error handling
func errs(errs ...error) (wasErr bool) {
	for _, err := range errs {
		if err != nil {
			fmt.Printf("> Err: %s\n", err)
			wasErr = true
		}
	}
	return
}