package pkg

import (
	"regexp"
	"strings"
)

func PhoneNormalize(phone string) (normalizedPhone string, err error) {
	// Удаление всех символов, кроме цифр
	re := regexp.MustCompile("[^0-9]+")
	normalized := re.ReplaceAllString(phone, "")

	// Если номер телефона начинается с "8", заменяем на "7" (для России)
	if strings.HasPrefix(normalized, "8") {
		normalized = "7" + normalized[1:]
	} else if strings.HasPrefix(normalized, "+7") {
		// Если номер начинается с "+7", удаляем "+"
		normalized = normalized[1:]
	}

	return normalized, nil
}
