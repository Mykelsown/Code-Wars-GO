package fiveKyu

func FirstNonRepeating(str string) string {

    if str == "" {
        return ""
    }

    // -------------------------
    // Step 1: Convert to lowercase (Unicode-safe)
    // -------------------------
    originalRunes := []rune(str)
    lowerRunes := make([]rune, len(originalRunes))

    for i, r := range originalRunes {
        if r >= 'A' && r <= 'Z' {
            lowerRunes[i] = r + ('a' - 'A')
        } else {
            lowerRunes[i] = r
        }
    }

    // -------------------------
    // Step 2: Check duplicates using nested loops
    // -------------------------
    elementIsDuplicate := make(map[rune]bool)

    for i, char1 := range lowerRunes {

        // Step 3: Avoid re-processing existing characters
        if _, exists := elementIsDuplicate[char1]; exists {
            continue
        }

        isDuplicate := false

        for j, char2 := range lowerRunes {
            if i != j && char1 == char2 {
                isDuplicate = true
                break
            }
        }

        elementIsDuplicate[char1] = isDuplicate
    }

    // -------------------------
    // Step 4: Find first non-repeating character
    // -------------------------
    for i, r := range lowerRunes {
        if elementIsDuplicate[r] == false {
            return string(originalRunes[i])
        }
    }

    // -------------------------
    // Step 5: If none found
    // -------------------------
    return ""
}