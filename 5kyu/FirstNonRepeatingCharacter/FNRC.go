package fiveKyu

import "strings"

func Fnrc(str string) string{

	// Conversion of the string to lowercase
    result := []rune(str)

    for i, r := range result {
        if r >= 'A' && r <= 'Z' {
            result[i] = r + ('a' - 'A')
        }
    }

    convertedStr := string(result)

}