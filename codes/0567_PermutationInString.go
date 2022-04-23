package algorithmn

func checkInclusion(s1 string, s2 string) bool {
    if len(s2) < len(s1) {
        return false
    }
    
    var freq1 [256]int
    var freq2 [256]int
    var length int = len(s1)
    
    for i := 0; i < len(s1); i++ {
        freq1[s1[i]-'a']++
    }
    
    for left, right := 0, 0; right < len(s2); right++ {
        freq2[s2[right]-'a']++
        if right - left >= length {
            freq2[s2[left]-'a']--
            left++
        }
        if assertEqual(freq1, freq2) {
            return true
        }
    }
    return false
}

func assertEqual(arr1 [256]int, arr2 [256]int) bool {
    for i := 0; i < len(arr1); i++ {
        if arr1[i] != arr2[i] {
            return false
        }
    }
    return true
}