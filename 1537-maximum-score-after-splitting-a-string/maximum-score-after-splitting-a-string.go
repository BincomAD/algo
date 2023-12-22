func maxScore(str string) int {
    result := 0
    for i := 0; i < len(str)-1; i++ {
        currentScore := 0
        for j := 0; j <= i; j++ {
            if str[j] == '0' {
                currentScore++
            }
        }
        for j := i + 1; j < len(str); j++ {
            if str[j] == '1' {
                currentScore++
            }
        }
        result = max(result, currentScore)
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}