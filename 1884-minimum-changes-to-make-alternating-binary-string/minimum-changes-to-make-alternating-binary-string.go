func minOperations(s string) int {
    first, second := 0, 0
    
    for i := 0; i < len(s); i++ {
        if i%2 == 0 {
            if s[i] != '0' {
                first++
            } else {
                second++
            }
        } else {
            if s[i] != '1' {
                first++
            } else {
                second++
            }
        }
    }
    
    return min(len(s)-first, len(s)-second)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    
    return b
}