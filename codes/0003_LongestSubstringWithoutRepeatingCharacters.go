package algorithmn

func lengthOfLongestSubstring(s string) int {
    max, chars := 0, make(map[byte]int)
    for left, right := 0, 0; right < len(s); right++ {
        index := s[right]
        if _, ok := chars[index]; ok {
            if chars[index] + 1 > left {
                left = chars[index] + 1
            }
        }
        
        chars[index] = right
        if right - left + 1 > max {
            max = right - left + 1
        }
    }
    return max
}