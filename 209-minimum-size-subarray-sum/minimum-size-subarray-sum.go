func minSubArrayLen(target int, nums []int) int {
    prefix := make([]int, len(nums) + 1)

    for i := 0; i < len(nums); i++ {
        prefix[i+1] = prefix[i] + nums[i]
    }

    i := 0
    j := 1

    minlen := math.MaxInt

    for j < len(prefix){
        diff := prefix[j] - prefix[i]
        if diff < target {
            j++
        } else {
            if j - i < minlen{
                minlen = j - i
            }
            i++
        }
    }
    
    if minlen == math.MaxInt {
        return 0
    }

    return minlen
}