package algorithmn

func lengthOfLongestSubstring(s string) int {
    max, words, chars := 0, []byte(s), make(map[string]int)
    for left, right := 0, 0; right < len(words); right++ {
        char := string(words[right])
        if _, ok := chars[char]; ok {
            if chars[char] + 1 > left {
                left = chars[char] + 1
            }
        }
        
        chars[char] = right
        if right - left + 1 > max {
            max = right - left + 1
        }
    }
    return max
}