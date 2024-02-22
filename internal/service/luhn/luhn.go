package luhn

import "strconv"

func IsValidLuhn(orderNumber string) bool {
	var sum int
	alternate := false

	for i := len(orderNumber) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(orderNumber[i]))

		if err != nil {
			return false
		}

		if alternate {
			digit *= 2

			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		alternate = !alternate
	}

	return sum%10 == 0
}
