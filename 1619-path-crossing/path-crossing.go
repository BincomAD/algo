func isPathCrossing(path string) bool {
    was := make(map[string]bool, len(path))

    x, y := 0, 0

    was[fmt.Sprintf("%s-%s", x, y)] = true

    for _, s := range path {

        if string(s) == "N" {
            y++
        } else if string(s) == "E" {
            x++
        } else if string(s) == "S" {
            y--
        } else {
            x--
        }

        if was[fmt.Sprintf("%s-%s", x, y)] {
            return true
        }

        was[fmt.Sprintf("%s-%s", x, y)] = true
    }

    return false
}