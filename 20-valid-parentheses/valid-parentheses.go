func isValid(s string) bool {
    p := map[rune]rune{
        '(':')',
        '[':']',
        '{':'}',
    }

    stack := []rune{}

    for i := 0; i < len(s); i++ {
        if _, ok := p[rune(s[i])]; ok {
            stack = append(stack, rune(s[i]))
        } else if len(stack) == 0 || p[stack[len(stack)-1]] != rune(s[i]) {
            return false
        } else {
            stack = stack[:len(stack)-1]
        }
    }

    return len(stack) == 0
}