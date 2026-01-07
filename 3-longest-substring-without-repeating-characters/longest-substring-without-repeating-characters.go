func lengthOfLongestSubstring(s string) int {
    left := 0
    maxLen := 0
    itemMap := make(map[byte]int) // یا map[rune]int برای یونیکد

    for right := 0; right < len(s); right++ {
        char := s[right]
        if idx, ok := itemMap[char]; ok && idx >= left {
            left = idx + 1
        }
        itemMap[char] = right
        if currentLen := right - left + 1; currentLen > maxLen {
            maxLen = currentLen
        }
    }

    return maxLen
}
