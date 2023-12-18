class Solution:
    def maxProductDifference(self, nums: List[int]) -> int:
        mas = sorted(nums)
        return(mas[-1]*mas[-2] - mas[0]*mas[1])