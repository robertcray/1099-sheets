//
// Robert Cray <rdcray@pm.me> 
// May 2023

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Deal with the annoying feature of go that dates in the
// Format m/d/year don't parse if the month or day are single digit
func pDate(dt string) (time.Time, error) {
	var val time.Time

	res1 := strings.Split(dt, "/")
	layout := "01/02/2006"
	a1, e := strconv.Atoi(res1[0])
	if e != nil {
		return val, e
	}
	a2, e := strconv.Atoi(res1[1])
	if e != nil {
		return val, e
	}
	a3, e := strconv.Atoi(res1[2])
	if e != nil {
		return val, e
	}
	s := fmt.Sprintf("%02d/%02d/%04d", a1, a2, a3)
	val, e = time.ParseInLocation(layout, s, time.UTC)
	if e != nil {
		return val, e
	}
	return val, nil
}

// Wrange is struct with start and end date for each week
type Wrange struct {
	dt1     time.Time
	dt2     time.Time
	minutes int
}

// Return next monday
func nextMonday(s time.Time) time.Time {
	var t time.Time
	wd := int(s.Weekday())
	monday := int(time.Monday)
	if wd == monday {
		t = s.AddDate(0, 0, 7)
	} else if wd < monday {
		t = s.AddDate(0, 0, monday-wd)
	} else {
		t = s.AddDate(0, 0, (6-wd)+2)
	}
	return t
}

// Check if date s is between d1 and d2
func inRange(s, d1, d2 time.Time) bool {
	if s.Before(d1) || s.After(d2) {
		return false
	}
	return true
}

func formatMin(n int) string {
	if n < 60 {
		return fmt.Sprintf("%d minutes", n)
	}
	h := n / 60
	m := n % 60
	return fmt.Sprintf("%d hours, %d minutes", h, m)
}

// Return the # of days in a month for a given year
func lastDayInMonth(year, month int) time.Time {
	fst := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	lst := fst.AddDate(0, 1, -1) // Add 1 month, subtract 1 day
	return lst
}

func weekRanges(year, month int) []Wrange {
	var wranges = []Wrange{}
	var start, end time.Time
	var dt2 time.Time

	start = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	last := lastDayInMonth(year, month)

	for {
		wr := new(Wrange)
		end = nextMonday(start)
		if end.After(last) {
			end = last
			dt2 = last
		} else {
			dt2 = end.AddDate(0, 0, -1)
		}
		wr.dt1 = start
		wr.dt2 = dt2
		wr.minutes = 0
		wranges = append(wranges, *wr)
		//fmt.Println(start, end)
		if dt2 == last {
			break
		}
		start = end
	}
	return wranges
}

// Parse the time, e.g. "15m", "1h", etc.
func parseT(s string) int {
	re1 := regexp.MustCompile("^([0-9]+)m$")
	re2 := regexp.MustCompile("^([0-9]+)hr*$")
	re3 := regexp.MustCompile("^([0-9]+)hr*([0-9]+)m$")
	if sm := re1.FindStringSubmatch(s); sm != nil {
		t, _ := strconv.Atoi(sm[1])
		return t
	}
	if sm := re2.FindStringSubmatch(s); sm != nil {
		t, _ := strconv.Atoi(sm[1])
		return t * 60
	}
	if sm := re3.FindStringSubmatch(s); sm != nil {
		t1, _ := strconv.Atoi(sm[1])
		t2, _ := strconv.Atoi(sm[2])
		return (t1 * 60) + t2
	}
	fmt.Println("Don't know what to do with ", s)
	os.Exit(1)
	return 0
}

func getData(wr []Wrange, credFile, spreadSheetID, tabName string) {
	var tspent int

	_ = tabName
	records := readSheet(credFile, spreadSheetID, tabName)
	for _, w := range records {
		dt1, e := pDate(w.dt)
		if e != nil {
			log.Fatal("Encountered bad date, exiting")
		}
		tspent = parseT(w.spent)
		for idx := 0; idx < len(wr); idx++ {
			if inRange(dt1, wr[idx].dt1, wr[idx].dt2) {
				wr[idx].minutes = wr[idx].minutes + tspent
			}
		}
	}
}

// Print each week with the data for that week
func printWR(wr []Wrange) int {
	days := [7]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	var val string

	total := 0
	for idx, a := range wr {
		if a.minutes == 0 {
			val = "--------"
		} else {
			val = formatMin(a.minutes)
		}
		fmt.Printf("Week %d - %s(%s) - %s(%s)     %s\n",
			idx+1,
			a.dt1.Format("01/02/2006"),
			days[a.dt1.Weekday()],
			a.dt2.Format("01/02/2006"),
			days[a.dt2.Weekday()], val)
		total = total + a.minutes
	}
	return total
}
