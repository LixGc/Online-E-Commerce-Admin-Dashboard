package helpers

import (
	"strconv"
	"strings"
)

// FormatRupiah formats an integer as Rupiah
func FormatRupiah(amount int) string {
	amountStr := strconv.FormatInt(int64(amount), 10)
	parts := strings.Split(amountStr, ".")
	integerPart := parts[0]
	formattedInteger := addCommas(integerPart)
	return "Rp. " + formattedInteger
}

func addCommas(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return addCommas(s[:n-3]) + "," + s[n-3:]
}
