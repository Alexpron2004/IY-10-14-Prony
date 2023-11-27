package pkg

func PhoneNormalize(phone string) (normalizedPhone string, err error) {
	return
}

package main

import (
 "fmt"
 "regexp"
 "strings"
)

func formatPhoneNumber(phoneNumber string) string {
 // Удаление всех символов, кроме цифр
 re := regexp.MustCompile("[^0-9]+")
 digitsOnly := re.ReplaceAllString(phoneNumber, "")

 // Приведение номера к единому формату
 if strings.HasPrefix(digitsOnly, "8") {
  digitsOnly = "7" + digitsOnly[1:]
 } else if strings.HasPrefix(digitsOnly, "7") {
  digitsOnly = "7" + digitsOnly[1:]
 } else if strings.HasPrefix(digitsOnly, "+7") {
  digitsOnly = "7" + digitsOnly[2:]
 }

 return digitsOnly
}

func main() {
 phoneNumber1 := "+7 (999) 123-45-67"
 phoneNumber2 := "8-999-123-45-67"

 formattedPhoneNumber1 := formatPhoneNumber(phoneNumber1)
 formattedPhoneNumber2 := formatPhoneNumber(phoneNumber2)

 fmt.Println(formattedPhoneNumber1)
 fmt.Println(formattedPhoneNumber2)
}
