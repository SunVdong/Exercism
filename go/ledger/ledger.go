package ledger

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type FormatType struct {
	Currency string
	Local    string
}

type EntryList []Entry

func (e Entry) FormatDate(locale string) (string, error) {
	t, err := time.Parse("2006-01-02", e.Date)
	if err != nil {
		return "", errors.New("invalid date")
	}

	if locale == "nl-NL" {
		return t.Format("02-01-2006"), nil
	} else if locale == "en-US" {
		return t.Format("01/02/2006"), nil
	}

	return "", errors.New("invalid locale")
}

func (e Entry) FormatDes() string {
	des := e.Description
	if len(des) > 25 {
		des = des[:22] + "..."
	}
	return des
}

func (e Entry) FormatChange(currency string, locale string) string {
	symbol := ""
	if currency == "EUR" {
		symbol = "€"
	}
	if currency == "USD" {
		symbol = "$"
	}
	isNegative := e.Change < 0

	//change := e.Change * -1
	//change_str := strconv.Itoa(change)
	//l := len(change_str)

	if locale == "nl-NL" {
		if float64(e.Change)/100 > 1000 {
			return fmt.Sprintf("%s %d.%3d,%2d ", symbol, e.Change/100/1000, e.Change/100%1000, e.Change%100)
		}
		if -1*e.Change > 100 {
			return fmt.Sprintf("%s %d,%d-", symbol, -1*e.Change/100, -1*e.Change%100)
		}
	}
	if isNegative {
		return fmt.Sprintf("(%s%.2f)", symbol, -1*float64(e.Change)/100)
	} else {
		return fmt.Sprintf(" %s%.2f ", symbol, float64(e.Change)/100)
	}

}

func (e Entry) StringWithFormat(currency string, locale string) (string, error) {
	t, err := e.FormatDate(locale)
	if err != nil {
		return "", errors.New("invalid date")
	}

	des := e.FormatDes()

	cha := e.FormatChange(currency, locale)

	return fmt.Sprintf("%-10s | %-25s | %13s\n", t, des, cha), nil
}

func (el EntryList) Len() int {
	return len(el)
}

// Less order : Date > Description > Change
func (el EntryList) Less(i, j int) bool {
	time_i, _ := time.Parse("2006-01-02", el[i].Date)
	time_j, _ := time.Parse("2006-01-02", el[j].Date)
	if time_i != time_j {
		return time_i.Before(time_j)
	}
	if el[i].Description != el[j].Description {
		return el[i].Description < el[j].Description
	}
	return el[i].Change < el[j].Change
}

func (el EntryList) Swap(i, j int) {
	el[i], el[j] = el[j], el[i]
}

var headers = map[string][]string{
	"nl-NL": {"Datum", "Omschrijving", "Verandering"},
	"en-US": {"Date", "Description", "Change"},
}

func formatHeader(locale string) string {
	titles, ok := headers[locale]
	if !ok {
		return "Unsupported locale\n"
	}

	return fmt.Sprintf("%-10s | %-25s | %s\n", titles[0], titles[1], titles[2])
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if currency != "EUR" && currency != "USD" {
		return "", errors.New("invalid currency")
	}
	if locale != "nl-NL" && locale != "en-US" {
		return "", errors.New("invalid locale")
	}
	sort.Sort(EntryList(entries))

	var res string
	res = formatHeader(locale)
	for _, e := range entries {
		s, err := e.StringWithFormat(currency, locale)
		if err != nil {
			return "", err
		}
		res += s
	}

	return res, nil

}

// func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
// 	if currency != "EUR" && currency != "USD" {
// 		return "", errors.New("invalid currency")
// 	}
// 	if locale != "nl-NL" && locale != "en-US" {
// 		return "", errors.New("invalid locale")
// 	}

// 	var entriesCopy []Entry
// 	entriesCopy = append(entriesCopy, entries...)
// 	m1 := map[bool]int{true: 0, false: 1}
// 	m2 := map[bool]int{true: -1, false: 1}
// 	es := entriesCopy
// 	for len(es) > 1 {
// 		first, rest := es[0], es[1:]
// 		success := false
// 		for !success {
// 			success = true
// 			for i, e := range rest {
// 				if (m1[e.Date == first.Date]*m2[e.Date < first.Date]*4 +
// 					m1[e.Description == first.Description]*m2[e.Description < first.Description]*2 +
// 					m1[e.Change == first.Change]*m2[e.Change < first.Change]*1) < 0 {
// 					es[0], es[i+1] = es[i+1], es[0]
// 					success = false
// 				}
// 			}
// 		}
// 		es = es[1:]
// 	}
// 	s := formatHeader(locale)

// 	// Parallelism, always a great idea
// 	co := make(chan struct {
// 		i int
// 		s string
// 		e error
// 	})
// 	for i, et := range entriesCopy {
// 		go func(i int, entry Entry) {
// 			t, err := time.Parse("2006-01-02", entry.Date)
// 			if err != nil {
// 				co <- struct {
// 					i int
// 					s string
// 					e error
// 				}{e: errors.New("")}
// 			}
// 			var d string
// 			if locale == "nl-NL" {
// 				d = t.Format("02-01-2006")
// 			} else if locale == "en-US" {
// 				d = t.Format("01/02/2006")
// 			}

// 			de := entry.Description
// 			if len(de) > 25 {
// 				de = de[:22] + "..."
// 			} else {
// 				de = de + strings.Repeat(" ", 25-len(de))
// 			}

// 			negative := false
// 			cents := entry.Change
// 			if cents < 0 {
// 				cents = cents * -1
// 				negative = true
// 			}
// 			var a string
// 			if locale == "nl-NL" {
// 				if currency == "EUR" {
// 					a += "€"
// 				} else if currency == "USD" {
// 					a += "$"
// 				} else {
// 					co <- struct {
// 						i int
// 						s string
// 						e error
// 					}{e: errors.New("")}
// 				}
// 				a += " "
// 				centsStr := strconv.Itoa(cents)
// 				switch len(centsStr) {
// 				case 1:
// 					centsStr = "00" + centsStr
// 				case 2:
// 					centsStr = "0" + centsStr
// 				}
// 				rest := centsStr[:len(centsStr)-2]
// 				var parts []string
// 				for len(rest) > 3 {
// 					parts = append(parts, rest[len(rest)-3:])
// 					rest = rest[:len(rest)-3]
// 				}
// 				if len(rest) > 0 {
// 					parts = append(parts, rest)
// 				}
// 				for i := len(parts) - 1; i >= 0; i-- {
// 					a += parts[i] + "."
// 				}
// 				a = a[:len(a)-1]
// 				a += ","
// 				a += centsStr[len(centsStr)-2:]
// 				if negative {
// 					a += "-"
// 				} else {
// 					a += " "
// 				}
// 			} else if locale == "en-US" {
// 				if negative {
// 					a += "("
// 				}
// 				if currency == "EUR" {
// 					a += "€"
// 				} else if currency == "USD" {
// 					a += "$"
// 				} else {
// 					co <- struct {
// 						i int
// 						s string
// 						e error
// 					}{e: errors.New("")}
// 				}
// 				centsStr := strconv.Itoa(cents)
// 				switch len(centsStr) {
// 				case 1:
// 					centsStr = "00" + centsStr
// 				case 2:
// 					centsStr = "0" + centsStr
// 				}
// 				rest := centsStr[:len(centsStr)-2]
// 				var parts []string
// 				for len(rest) > 3 {
// 					parts = append(parts, rest[len(rest)-3:])
// 					rest = rest[:len(rest)-3]
// 				}
// 				if len(rest) > 0 {
// 					parts = append(parts, rest)
// 				}
// 				for i := len(parts) - 1; i >= 0; i-- {
// 					a += parts[i] + ","
// 				}
// 				a = a[:len(a)-1]
// 				a += "."
// 				a += centsStr[len(centsStr)-2:]
// 				if negative {
// 					a += ")"
// 				} else {
// 					a += " "
// 				}
// 			} else {
// 				co <- struct {
// 					i int
// 					s string
// 					e error
// 				}{e: errors.New("")}
// 			}
// 			var al int
// 			for range a {
// 				al++
// 			}
// 			co <- struct {
// 				i int
// 				s string
// 				e error
// 			}{i: i, s: d + strings.Repeat(" ", 10-len(d)) + " | " + de + " | " +
// 				strings.Repeat(" ", 13-al) + a + "\n"}
// 		}(i, et)
// 	}
// 	ss := make([]string, len(entriesCopy))
// 	for range entriesCopy {
// 		v := <-co
// 		if v.e != nil {
// 			return "", v.e
// 		}
// 		ss[v.i] = v.s
// 	}
// 	for i := 0; i < len(entriesCopy); i++ {
// 		s += ss[i]
// 	}
// 	return s, nil
// }
