func minSubArrayLen(target int, nums []int) int {
    prefixsum := make([]int, len(nums)+1)

    for i := 1; i < len(nums) + 1; i++ {
        prefixsum[i] = prefixsum[i-1] + nums[i-1]
    }

    i, j := 0, 1
    max := math.MaxInt
    flag := false
    for i < len(prefixsum) {
        if j < len(nums) && prefixsum[j] - prefixsum[i] < target {
            j++
        } else if prefixsum[j] - prefixsum[i] >= target {
            if max > j - i {
                flag = true
                max = j - i
            }
            i++
        } else {
            i++
        }
    }

    if flag {
        return max
    } else {
        return 0
    }

}