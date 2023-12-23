func twoSum(nums []int, target int) []int {
    ch := make(map[int]int)
    res := []int{}
    was := make(map[int]bool)
    for i := 0; i < len(nums); i++ {
        ch[nums[i]] = i        
    }

    for i := 0; i < len(nums); i++ {
        if num, ok := ch[target-nums[i]]; ok {
            if !was[i] && !was[num] && num != i {
                res = append(res, i)
                res = append(res, num)
                was[i] = true
                was[num] = true
            }
        }
    }

    return res
}