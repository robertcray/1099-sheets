//
// Robert Cray <rdcray@pm.me>
// May 2023
//
// Read google sheet for 1099 work and create summary for input into invoice
//

package main

import (
	"fmt"
)

func main() {
	opts := getargs()
	credFile, spreadSheetID, rate, tabName := parseToml(opts.config)
	if opts.rate != 0.0 {
		rate = opts.rate
	}

	wr := weekRanges(opts.year, opts.month)
	getData(wr, credFile, spreadSheetID, tabName)
	tot := printWR(wr)
	fmt.Printf("\nTotal: %s\n", formatMin(tot))
	h := tot / 60
	m := tot % 60

	invoice := (rate * float64(h)) + ((float64(m) / 60.0) * rate)
	fmt.Printf("Total rate: $%.2f (at $%.2f/hr)\n", invoice, rate)
}
