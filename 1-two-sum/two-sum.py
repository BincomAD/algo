class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        t = {}

        for i, ch in enumerate(nums):
            t[ch] = i
        
        for stan in range(len(nums)):
            tar = target - nums[stan]
            if tar in t and t[tar] != stan:
                return [t[tar], stan]

