func intToRoman(num int) string {
    t := map[int]string{
        1: "I",
        4: "IV",
        5: "V",
        9: "IX",
        10: "X",
        40: "XL",
        50: "L",
        90: "XC",
        100: "C",
        400: "CD",
        500: "D",
        900: "CM",
        1000: "M",
    }

    arr := []int{1,4,5,9,10,40,50,90,100,400,500,900,1000}
    i := len(arr) - 1
    res := ""
    for num > 0 {
        if num < arr[i] {
            i--
        } else {
            num -= arr[i]
            res = res + t[arr[i]]
        }
    }

    return res
}