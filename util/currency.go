package util

const (
	INR = "INR"
	EUR = "EUR"
	YEN = "YEN"
)

func IsValidCurrency(currency string) bool {
	switch currency {
	case INR, EUR, YEN:
		return true
	}
	return false
}