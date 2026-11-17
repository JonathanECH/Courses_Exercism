package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

/*
Refactoring Changelog:
1. Early Validation: Added `validateInput` to validate `currency` and `locale` immediately, eliminating the bizarre recursive call with dummy data.
2. Simplified Sorting: Replaced the obfuscated custom sort (which used boolean maps multiplied by 4, 2, and 1) with Go's standard `sort.Slice`.
3. Header Simplification: Replaced 18 lines of `strings.Repeat` calculations with simple string literals based on the locale.
4. Concurrency Removal: Eliminated unnecessary goroutines and channels (`// Parallelism, always a great idea`) which added significant overhead and complexity to an in-memory formatting task.
5. Modular Currency Formatting: Extracted `formatAmount` to cleanly handle currency symbols, negative formats (parentheses vs minus), thousands/decimal separators, and 13-character alignment using `utf8.RuneCountInString`.
6. Date & Description Formatting: Replaced manual character splitting and string concatenation with readable slicing and standard `fmt.Sprintf` formatting.
*/

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

// validateInput checks that the currency and locale are supported before proceeding.
func validateInput(currency, locale string) bool {
	return (currency == "USD" || currency == "EUR") && (locale == "en-US" || locale == "nl-NL")
}

// formatAmount formats an amount in cents according to the currency and locale rules,
// padded to a visual width of 13 characters.
func formatAmount(change int, currency, locale string) string {
	symbol := "$"
	if currency == "EUR" {
		symbol = "€"
	}

	isNegative := change < 0
	absChange := change
	if isNegative {
		absChange = -change
	}

	// Determine thousands and decimal separators by locale
	var thousandsSep, decimalSep string
	if locale == "nl-NL" {
		thousandsSep, decimalSep = ".", ","
	} else {
		thousandsSep, decimalSep = ",", "."
	}

	cents, units := absChange%100, absChange/100

	// Format thousands grouping from right to left
	u := strconv.Itoa(units)
	var parts []string
	for len(u) > 3 {
		parts = append([]string{u[len(u)-3:]}, parts...)
		u = u[:len(u)-3]
	}
	parts = append([]string{u}, parts...)
	formattedUnits := strings.Join(parts, thousandsSep)

	// Apply locale-specific negative number formatting
	var priceFormt string
	if locale == "nl-NL" {
		if !isNegative {
			priceFormt = fmt.Sprintf("%s %s%s%02d ", symbol, formattedUnits, decimalSep, cents)
		} else {
			priceFormt = fmt.Sprintf("%s -%s%s%02d ", symbol, formattedUnits, decimalSep, cents)
		}
	} else {
		if !isNegative {
			priceFormt = fmt.Sprintf("%s%s%s%02d ", symbol, formattedUnits, decimalSep, cents)
		} else {
			priceFormt = fmt.Sprintf("(%s%s%s%02d)", symbol, formattedUnits, decimalSep, cents)
		}
	}

	// Right-align to 13 runes (counting visual characters, not UTF-8 bytes for €)
	espacios := strings.Repeat(" ", 13-utf8.RuneCountInString(priceFormt))
	return espacios + priceFormt
}

// FormatLedger formats the full ledger output with localized header, sorted entries, and columns.
func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	// 1. Validate inputs upfront
	if !validateInput(currency, locale) {
		return "", errors.New("The data are incorrect!")
	}

	// Copy slice to avoid mutating the original input
	entriesCopy := append([]Entry(nil), entries...)

	// 2. Sort entries by Date, then Description, then Change
	sort.Slice(entriesCopy, func(i, j int) bool {
		a, b := entriesCopy[i], entriesCopy[j]
		if a.Date != b.Date {
			return a.Date < b.Date
		}
		if a.Description != b.Description {
			return a.Description < b.Description
		}
		return a.Change < b.Change
	})

	// 3. Select table header by locale
	var header string
	if locale == "nl-NL" {
		header = "Datum      | Omschrijving              | Verandering  \n"
	} else {
		header = "Date       | Description               | Change       \n"
	}

	// 4. Format each entry row
	for _, entry := range entriesCopy {
		date, desc, change := entry.Date, entry.Description, entry.Change

		// Validate date format (YYYY-MM-DD)
		if len(date) != 10 || date[4] != '-' || date[7] != '-' {
			return "", errors.New("Invalid date")
		}

		day, month, year := date[8:10], date[5:7], date[0:4]
		var dateNew string
		if locale == "nl-NL" {
			dateNew = fmt.Sprintf("%s-%s-%s", day, month, year)
		} else {
			dateNew = fmt.Sprintf("%s/%s/%s", month, day, year)
		}

		// Truncate or pad description to exactly 25 characters
		if len(desc) > 25 {
			desc = desc[:22] + "..."
		} else {
			desc = fmt.Sprintf("%-25s", desc)
		}

		// Format amount and append complete table row
		amount := formatAmount(change, currency, locale)
		header += fmt.Sprintf("%s | %s | %s\n", dateNew, desc, amount)
	}

	return header, nil
}
