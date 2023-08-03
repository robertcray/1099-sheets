//
// Robert Cray <rdcray@pm.me> 
// May 2023

package main

import (
	"flag"
	"time"
)

type options struct {
	month  int
	year   int
	rate   float64
	config string
}

// Get args and return in struct
// month defaults to last month
// year defaults to this year
// Rate defaults to 0.0 and will typically be retrieved from the 1099.toml
// file, but if a non-zero --rate option is given it will override 1099.toml
func getargs() options {
	var opts options
	flag.StringVar(&opts.config, "config", "/usr/local/etc/1099.toml", "Config")
	flag.IntVar(&opts.month, "month", 0, "Month")
	flag.IntVar(&opts.year, "year", 0, "year")
	flag.Float64Var(&opts.rate, "rate", 0.0, "Rate")
	flag.Parse()

	now := time.Now()
	if opts.year == 0 {
		opts.year = now.Year()
	}
	if opts.month == 0 {
		month := int(now.Month())
		if month > 1 {
			opts.month = month - 1
		} else {
			opts.month = 12
			opts.year = opts.year - 1
		}
	}
	return opts
}
