package util

// Setup the constants
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// Setup the function to check if the currency is valid
func IsSupportedCurrency(currency string) bool {
	switch currency {
		case USD, EUR, CAD:
			return true
	}
	return false
}