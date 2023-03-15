package util

const (
	USD = "USD"
	EUR = "EUR"
	GBP = "GBP"
	INR = "INR"
	CAD = "CAD"
	AUD = "AUD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, INR, GBP, CAD, AUD:
		return true
	default:
		return false
	}
}
