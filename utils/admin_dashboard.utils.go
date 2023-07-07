package utils

import "errors"

func ParseMonth(query string) (string, error) {
	months := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}

	month, ok := months[query]
	if !ok {
		return "", errors.New("Invalid month")
	}

	return month, nil
}
