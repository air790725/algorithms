package algorithmn

func reverseWords(s string) string {
    i, j, str := 0, 0, []byte(s + " ")
    for i < len(str) {
        if string(str[i:i+1]) == " " {
            for left, right := j, i - 1; left < right; {
                str[left], str[right] = str[right], str[left]
                left++
                right--
            }
            j = i+1
        }
        i++
    }
    return string(str[:len(s)])
}