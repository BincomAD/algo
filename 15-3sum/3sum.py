class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        answ = set()
        dups = set()

        def two_sum(i):
            target = 0 - nums[i]
            seen = set()
            for j in range(len(nums)):
                if j != i:
                    compl = target - nums[j]
                    if compl in seen:
                        answ.add(tuple(sorted([nums[j], nums[i], compl])))
                    seen.add(nums[j])
        
        for i in range(len(nums)):
            if not nums[i] in dups:
                two_sum(i)
            dups.add(nums[i])
        
        return answ


