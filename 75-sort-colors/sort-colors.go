func sortColors(nums []int)  {
    r, w, b := 0, 0, 0

    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            r++
        } else if nums[i] == 1 {
            w++
        } else {
            b++
        }
    }

    w += r
    b += w

    for i := 0; i < len(nums); i++ {
        if i < r {
            nums[i] = 0
        } else if i < w {
            nums[i] = 1
        } else {
            nums[i] = 2
        }
    }
}