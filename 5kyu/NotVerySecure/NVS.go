package fiveKyu

func Alphanumeric(str string) bool {
	 // Step-1: Check for empty string
    if len(str) < 1 {
        return false
    }

    // Step-2: Loop through each character as rune
    for _, char := range str {

        // Step-3: Check ASCII ranges
        if (char >= 'A' && char <= 'Z') ||
           (char >= 'a' && char <= 'z') ||
           (char >= '0' && char <= '9') {

            // valid character → continue checking
            continue

        } else {
            // invalid character found
            return false
        }
    }

    // If loop completes with no invalid characters
    return true
}
