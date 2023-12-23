func distributeCandies(candyType []int) int {
    candy := make(map[int]bool)
    count := 0

    for i := 0; i < len(candyType); i++ {
        if candy[candyType[i]] {

        } else {
            count++
            candy[candyType[i]] = true
        }
    }

    if count > len(candyType)/2 {
        return len(candyType)/2 
    } else {
        return count
    }
}