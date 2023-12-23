func romanToInt(s string) int {
    t := map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }

    sum := 0
    for i := 0; i < len(s)-1; i++ {
        if t[string(s[i])] < t[string(s[i+1])] {
            sum -= t[string(s[i])]
        } else {
            sum += t[string(s[i])]
        }
    }
    sum += t[string(s[len(s)-1])]

    return sum
}